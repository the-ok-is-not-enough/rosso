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

func Get() *Request {
   ref := new(url.URL)
   return New_Request(http.MethodGet, ref)
}

func Get_Parse(s string) (*Request, error) {
   ref, err := url.Parse(s)
   if err != nil {
      return nil, err
   }
   return New_Request(http.MethodGet, ref), nil
}

func New_Request(method string, ref *url.URL) *Request {
   var r Request
   r.Request = new(http.Request) // .Request
   r.Header = make(http.Header) // .Request.Header
   r.Method = method // .Request.Method
   r.ProtoMajor = 1 // .Request.ProtoMajor
   r.ProtoMinor = 1 // .Request.ProtoMinor
   r.URL = ref
   return &r
}

func Post() *Request {
   ref := new(url.URL)
   return New_Request(http.MethodPost, ref)
}

func Post_Parse(s string) (*Request, error) {
   ref, err := url.Parse(s)
   if err != nil {
      return nil, err
   }
   return New_Request(http.MethodPost, ref), nil
}

func (r Request) Body_Bytes(b []byte) {
   read := bytes.NewReader(b)
   r.Body = io.NopCloser(read)
}

func (r Request) Body_String(s string) {
   read := strings.NewReader(s)
   r.Body = io.NopCloser(read)
}
