package http

import (
   "bytes"
   "io"
   "net/http"
   "net/url"
   "strings"
)

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

func Get() Request {
   return New_Request("GET")
}

func Post() Request {
   return New_Request("POST")
}

type Request struct {
   *http.Request
}

func New_Request(method string) Request {
   var r Request
   r.Request = new(http.Request) // .Request
   r.Header = make(http.Header) // .Request.Header
   r.Method = method // .Request.Method
   r.ProtoMajor = 1 // .Request.ProtoMajor
   r.ProtoMinor = 1 // .Request.ProtoMinor
   r.URL = new(url.URL) // .Request.URL
   return r
}
