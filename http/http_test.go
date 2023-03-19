package http

import (
   "net/http"
   "os"
   "strings"
   "testing"
)

func Test_Request(t *testing.T) {
   req := New_Request()
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
   req := New_Request()
   err := req.Set_URL("http://httpbin.org/get")
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
   req := New_Request()
   req.Method = "POST"
   req.URL.Scheme = "http"
   req.URL.Host = "httpbin.org"
   req.URL.Path = "/post"
   req.Set_Body(strings.NewReader("hello=world"))
   res, err := new(http.Transport).RoundTrip(req.Request)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
func Test_Client_Get(t *testing.T) {
   res, err := Default_Client.Status(302).Get("http://godocs.io")
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
}

func Test_Client_Do(t *testing.T) {
   req, err := http.NewRequest("GET", "http://godocs.io", nil)
   if err != nil {
      t.Fatal(err)
   }
   res, err := Default_Client.Status(302).Do(req)
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
}
