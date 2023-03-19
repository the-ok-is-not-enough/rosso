package main

import (
   "2a.pages.dev/rosso/strconv"
   "bytes"
   "embed"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "text/template"
)

//go:embed _template.go
var content embed.FS

func write_request(req *http.Request, dst io.Writer) error {
   var v values
   if req.Body != nil && req.Method != "GET" {
      body, err := io.ReadAll(req.Body)
      if err != nil {
         return err
      }
      req.Body = io.NopCloser(bytes.NewReader(body))
      v.Raw_Req_Body = strconv.Quote(string(body))
      v.Req_Body = "io.NopCloser(req_body)"
   } else {
      v.Raw_Req_Body = `""`
      v.Req_Body = "nil"
   }
   v.Query = req.URL.Query()
   v.Request = req
   temp, err := template.ParseFS(content, "_template.go")
   if err != nil {
      return err
   }
   return temp.Execute(dst, v)
}

type values struct {
   *http.Request
   Query url.Values
   Req_Body string
   Raw_Req_Body string
}

func write(req *http.Request, file *os.File) error {
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if file == os.Stdout {
      dump, err := httputil.DumpResponse(res, true)
      if err != nil {
         return err
      }
      if !strconv.Valid(dump) {
         dump = strconv.AppendQuote(nil, string(dump))
      }
      file.Write(dump)
   } else {
      dump, err := httputil.DumpResponse(res, false)
      if err != nil {
         return err
      }
      os.Stdout.Write(dump)
      if _, err := file.ReadFrom(res.Body); err != nil {
         return err
      }
   }
   return nil
}

type flags struct {
   golang bool
   https bool
   name string
   output string
}
