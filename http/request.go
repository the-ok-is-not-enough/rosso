package http

import (
   "io"
   "net/http"
   "net/url"
)

type Request struct {
   *http.Request
}

func New_Request() Request {
   var r Request
   // first
   r.Request = new(http.Request)
   // second
   r.Header = make(http.Header)
   r.Method = "GET"
   r.URL = new(url.URL)
   r.URL.Scheme = "http"
   return r
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
