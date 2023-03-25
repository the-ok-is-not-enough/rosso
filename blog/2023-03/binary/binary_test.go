package quotedprintable

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "testing"
)

func Test_Binary(t *testing.T) {
   res, err := http.Get(binary)
   if err != nil {
      t.Fatal(err)
   }
   dump, err := httputil.DumpResponse(res, true)
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
   fmt.Println(string(escape(dump)))
}

const (
   binary = "https://picsum.photos/1"
   text = "http://httpbin.org/get"
)

func Test_Text(t *testing.T) {
   res, err := http.Get(text)
   if err != nil {
      t.Fatal(err)
   }
   dump, err := httputil.DumpResponse(res, true)
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
   fmt.Println(string(escape(dump)))
}
