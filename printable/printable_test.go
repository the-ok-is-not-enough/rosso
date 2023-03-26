package printable

import (
   "bytes"
   "net/http"
   "net/http/httputil"
   "testing"
)

func Test_Binary(t *testing.T) {
   src, dst, err := round_trip(binary)
   if err != nil {
      t.Fatal(err)
   }
   if !bytes.Equal(src, dst) {
      t.Fatal(dst)
   }
}

func round_trip(s string) ([]byte, []byte, error) {
   res, err := http.Get(s)
   if err != nil {
      return nil, nil, err
   }
   src, err := httputil.DumpResponse(res, true)
   if err != nil {
      return nil, nil, err
   }
   if err := res.Body.Close(); err != nil {
      return nil, nil, err
   }
   dst, err := decode(Encode(src))
   if err != nil {
      return nil, nil, err
   }
   return src, dst, nil
}

const (
   binary = "https://picsum.photos/1"
   text = "http://httpbin.org/get"
)

func Test_Text(t *testing.T) {
   src, dst, err := round_trip(text)
   if err != nil {
      t.Fatal(err)
   }
   if !bytes.Equal(src, dst) {
      t.Fatal(dst)
   }
}
func Test_Append(t *testing.T) {
   var b []byte
   b = New_Number(123).Cardinal(nil)
   if s := string(b); s != "123" {
      t.Fatal(s)
   }
   b = New_Number(1234).Cardinal(nil)
   if s := string(b); s != "1.23 thousand" {
      t.Fatal(s)
   }
   b = New_Number(123).Size(nil)
   if s := string(b); s != "123 byte" {
      t.Fatal(s)
   }
   b = New_Number(1234).Size(nil)
   if s := string(b); s != "1.23 kilobyte" {
      t.Fatal(s)
   }
   b = Ratio(1234, 10).Rate(nil)
   if s := string(b); s != "123 byte/s" {
      t.Fatal(s)
   }
   b = Ratio(12345, 10).Rate(nil)
   if s := string(b); s != "1.23 kilobyte/s" {
      t.Fatal(s)
   }
   b = Ratio(1234, 10000).Percent(nil)
   if s := string(b); s != "12.34%" {
      t.Fatal(s)
   }
}
