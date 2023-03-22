package tls

import (
   "fmt"
   "net/http"
   "testing"
)

var tests = []struct {
   in string
   out string
} {
   {Android_API_24, "8fcaa9e4a15f48af0a7d396e3fa5c5eb"},
   {Android_API_25, "9fc6ef6efc99b933c5e2d8fcf4f68955"},
   {Android_API_26, "d8c87b9bfde38897979e41242626c2f3"},
   {Android_API_29, "9b02ebd3a43b62d825e1ac605b621dc8"},
}

func Test_Fingerprint(t *testing.T) {
   for _, test := range tests {
      out := Fingerprint(test.in)
      if out != test.out {
         t.Fatal(out)
      }
   }
}

func Test_Transport(t *testing.T) {
   req, err := http.NewRequest("HEAD", "https://example.com", nil)
   if err != nil {
      t.Fatal(err)
   }
   hello, err := Parse(Android_API)
   if err != nil {
      t.Fatal(err)
   }
   res, err := hello.Transport().RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", res)
}
