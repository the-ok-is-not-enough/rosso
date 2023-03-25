package strconv

import (
   "fmt"
   "net/http"
   "net/http/httputil"
   "testing"
   "time"
   "unicode/utf8"
)

var refs = []string{
   "http://httpbin.org/get",
   "https://picsum.photos/1",
}

func Test_Quote(t *testing.T) {
   for _, ref := range refs {
      res, err := http.Get(ref)
      if err != nil {
         t.Fatal(err)
      }
      dump, err := httputil.DumpResponse(res, true)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", dump)
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
   }
}
