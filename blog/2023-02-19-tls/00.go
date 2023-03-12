package tls

import (
   "encoding/hex"
   "fmt"
   "net/http"
   "net/http/httputil"
   "net/url"
   "strings"
   "testing"
   "time"
)

func main() {
   req_body := url.Values{
      "Email": {email},
      "Passwd": {password},
      "client_sig": {""},
      "droidguard_results": {"-"},
   }.Encode()
   var hello Client_Hello
   err := hello.UnmarshalText(Android_API())
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/auth", strings.NewReader(req_body),
   )
   if err != nil {
      panic(err)
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := hello.Transport().RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res.Body, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}
