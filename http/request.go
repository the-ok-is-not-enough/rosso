package http

import (
   "io"
   "net/http"
   "net/url"
)

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

type Request struct {
   *http.Request
}

func (r Request) Set_Body(body io.Reader) {
   var ok bool
   r.Body, ok = body.(io.ReadCloser)
   if !ok {
      r.Body = io.NopCloser(body)
   }
}

func (r Request) Set_URL(ref string) error {
   var err error
   r.URL, err = url.Parse(ref)
   if err != nil {
      return err
   }
   return nil
}
