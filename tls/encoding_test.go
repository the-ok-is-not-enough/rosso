package tls

import (
   "bytes"
   "encoding/hex"
   "fmt"
   "net/http"
   "net/url"
   "strings"
   "testing"
   "time"
)

const (
   android = iota
   cURL
)

var handshakes = []string{
   android: "16030100bb010000b703034420d198e7852decbc117dc7f90550b98f2d643c95" +
   "4bf3361ddaf127ff921b04000024c02bc02ccca9c02fc030cca8009e009fc009" +
   "c00ac013c01400330039009c009d002f00350100006aff010001000000002200" +
   "2000001d636c69656e7473657276696365732e676f6f676c65617069732e636f" +
   "6d0017000000230000000d001600140601060305010503040104030301030302" +
   "0102030010000b000908687474702f312e31000b00020100000a000400020017",
   cURL: "1603010200010001fc03033356ee099c006213ecb9f7493ef981dd513761eae2" +
   "7eff36a177ebd353fc207520fa9ef53871b81af022e38d46ca9268be95889d6e964db8" +
   "18768ec86a68c7216f003e130213031301c02cc030009fcca9cca8ccaac02bc02f009e" +
   "c024c028006bc023c0270067c00ac0140039c009c0130033009d009c003d003c003500" +
   "2f00ff0100017500000010000e00000b6578616d706c652e636f6d000b000403000102" +
   "000a000c000a001d0017001e00190018337400000010000e000c02683208687474702f" +
   "312e31001600000017000000310000000d0030002e040305030603080708080809080a" +
   "080b080408050806040105010601030302030301020103020202040205020602002b00" +
   "09080304030303020301002d00020101003300260024001d002034107e2fb61cbfc3c8" +
   "27b3d574b57d9d5f5294bedb7ee350407c05d1a9396b5b001500b20000000000000000" +
   "0000000000000000000000000000000000000000000000000000000000000000000000" +
   "0000000000000000000000000000000000000000000000000000000000000000000000" +
   "0000000000000000000000000000000000000000000000000000000000000000000000" +
   "0000000000000000000000000000000000000000000000000000000000000000000000" +
   "000000000000000000000000000000000000000000000000000000000000",
}

func Test_UnmarshalBinary(t *testing.T) {
   for _, handshake := range handshakes {
      data, err := hex.DecodeString(handshake)
      if err != nil {
         t.Fatal(err)
      }
      hello := New_Client_Hello()
      if err := hello.UnmarshalBinary(data); err != nil {
         t.Fatal(err)
      }
      text, err := hello.MarshalText()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(string(text))
   }
}

func Test_UnmarshalText(t *testing.T) {
   body := url.Values{
      "Email": {email},
      "Passwd": {password},
      "client_sig": {""},
      "droidguard_results": {"-"},
   }.Encode()
   for _, test := range tests {
      hello := New_Client_Hello()
      err := hello.UnmarshalText([]byte(test.in))
      if err != nil {
         t.Fatal(err)
      }
      req, err := http.NewRequest(
         "POST", "https://android.googleapis.com/auth",
         strings.NewReader(body),
      )
      if err != nil {
         t.Fatal(err)
      }
      req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
      res, err := hello.Transport().RoundTrip(req)
      if err != nil {
         t.Fatal(err)
      }
      defer res.Body.Close()
      fmt.Println(res.Status, test.in)
      time.Sleep(time.Second)
   }
}

func Test_MarshalText(t *testing.T) {
   a := []byte(Android_API)
   hello := New_Client_Hello()
   err := hello.UnmarshalText(a)
   if err != nil {
      t.Fatal(err)
   }
   b, err := hello.MarshalText()
   if err != nil {
      t.Fatal(err)
   }
   if !bytes.Equal(b, a) {
      t.Fatal(b)
   }
}
