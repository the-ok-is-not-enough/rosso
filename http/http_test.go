package http

import (
   "net/http"
   "os"
   "testing"
)

func Test_Request(t *testing.T) {
   req := Get()
   req.URL.Scheme = "http"
   req.URL.Host = "httpbin.org"
   req.URL.Path = "/get"
   res, err := new(http.Transport).RoundTrip(req.Request)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}

func Test_URL(t *testing.T) {
   req := Get()
   err := req.URL_String("http://httpbin.org/get")
   if err != nil {
      t.Fatal(err)
   }
   res, err := new(http.Transport).RoundTrip(req.Request)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}

func Test_Body(t *testing.T) {
   req := Post()
   req.URL.Scheme = "http"
   req.URL.Host = "httpbin.org"
   req.URL.Path = "/post"
   req.Body_String("hello=world")
   res, err := new(http.Transport).RoundTrip(req.Request)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
