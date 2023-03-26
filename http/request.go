package http

import (
   "bytes"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type Request struct {
   *http.Request
}

func Get() Request {
   var r Request
   r.Request = new(http.Request) // .Request
   r.Header = make(http.Header) // .Request.Header
   r.Method = "GET" // .Request.Method
   r.URL = new(url.URL) // .Request.URL
   return r
}

func Post() Request {
   var r Request
   r.Request = new(http.Request) // .Request
   r.Header = make(http.Header) // .Request.Header
   r.Method = "POST" // .Request.Method
   r.URL = new(url.URL) // .Request.URL
   return r
}

func (r Request) Body_Bytes(b []byte) {
   read := bytes.NewReader(b)
   r.Body = io.NopCloser(read)
}

func (r Request) Body_String(s string) {
   read := strings.NewReader(s)
   r.Body = io.NopCloser(read)
}

func (r Request) URL_String(s string) error {
   var err error
   r.URL, err = url.Parse(s)
   if err != nil {
      return err
   }
   return nil
}
