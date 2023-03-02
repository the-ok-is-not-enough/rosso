package tls

import (
   "fmt"
   "net/http"
   "testing"
)

func Test_Transport(t *testing.T) {
   req, err := http.NewRequest("HEAD", "https://example.com", nil)
   if err != nil {
      t.Fatal(err)
   }
   hello := New_Client_Hello()
   if err := hello.UnmarshalText(Android_API()); err != nil {
      t.Fatal(err)
   }
   res, err := hello.Transport().RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", res)
}
