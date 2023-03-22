package http

import (
   "net/http"
   "testing"
)

func Test_Client(t *testing.T) {
   c := Default_Client.Clone()
   c.Log_Level = 9
   if Default_Client.Log_Level == 9 {
      t.Fatal("Level")
   }
   c.CheckRedirect = nil
   if Default_Client.CheckRedirect == nil {
      t.Fatal("Redirect")
   }
   c.Status = 9
   if Default_Client.Status == 9 {
      t.Fatal("Status")
   }
   c.Transport = new(http.Transport)
   if Default_Client.Transport != nil {
      t.Fatal("Transport")
   }
   req := Get()
   req.URL.Host = "godocs.io"
   Default_Client.Status = 302
   res, err := Default_Client.Do(req)
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
}
