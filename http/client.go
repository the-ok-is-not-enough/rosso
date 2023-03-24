package http

import (
   "2a.pages.dev/rosso/strconv"
   "bytes"
   "errors"
   "net/http"
   "net/http/httputil"
   "os"
)

type Client struct {
   Log_Level int // this needs to work with flag.IntVar
   Status int
   http.Client
}

var Default_Client = Client{
   Client: http.Client{
      CheckRedirect: func(*http.Request, []*http.Request) error {
         return http.ErrUseLastResponse
      },
   },
   Log_Level: 1,
   Status: http.StatusOK,
}

func (c Client) Clone() Client {
   return c
}

func (c Client) Do(req Request) (*Response, error) {
   switch c.Log_Level {
   case 1:
      os.Stderr.WriteString(req.Method)
      os.Stderr.WriteString(" ")
      os.Stderr.WriteString(req.URL.String())
      os.Stderr.WriteString("\n")
   case 2:
      dump, err := httputil.DumpRequest(req.Request, true)
      if err != nil {
         return nil, err
      }
      if !strconv.Valid(dump) {
         dump = strconv.AppendQuote(nil, string(dump))
      }
      if !bytes.HasSuffix(dump, []byte{'\n'}) {
         dump = append(dump, '\n')
      }
      os.Stderr.Write(dump)
   }
   res, err := c.Client.Do(req.Request)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != c.Status {
      return nil, errors.New(res.Status)
   }
   return res, nil
}
