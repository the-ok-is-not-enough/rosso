package main

import (
   "2a.pages.dev/rosso/http"
   "2a.pages.dev/rosso/strconv"
   "bufio"
   "bytes"
   "flag"
   "io"
   "net/http/httputil"
   "net/textproto"
   "net/url"
   "os"
   "strings"
   "text/template"
)

func read_request(in io.Reader) (*http.Request, error) {
   var req http.Request
   text := textproto.NewReader(bufio.NewReader(in))
   // .Method
   raw_method_path, err := text.ReadLine()
   if err != nil {
      return nil, err
   }
   method_path := strings.Fields(raw_method_path)
   req.Method = method_path[0]
   // .URL
   ref, err := url.Parse(method_path[1])
   if err != nil {
      return nil, err
   }
   req.URL = ref
   // .URL.Host
   head, err := text.ReadMIMEHeader()
   if err != nil {
      return nil, err
   }
   if req.URL.Host == "" {
      req.URL.Host = head.Get("Host")
   }
   // .Header
   req.Header = http.Header(head)
   // .Body
   buf := new(bytes.Buffer)
   length, err := text.R.WriteTo(buf)
   if err != nil {
      return nil, err
   }
   if length >= 1 {
      req.Body = io.NopCloser(buf)
   }
   req.ContentLength = length
   return &req, nil
}

const raw_temp = `package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   {{ range $key, $val := .Header -}}
      req.Header[{{ printf "%q" $key }}] = {{ printf "%#v" $val }}
   {{ end -}}
   req.Method = {{ printf "%q" .Method }}
   req.URL = new(url.URL)
   req.URL.Host = {{ printf "%q" .URL.Host }}
   req.URL.Path = {{ printf "%q" .URL.Path }}
   req.URL.RawPath = {{ printf "%q" .URL.RawPath }}
   val := make(url.Values)
   {{ range $key, $val := .Query -}}
      val[{{ printf "%q" $key }}] = {{ printf "%#v" $val }}
   {{ end -}}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = {{ printf "%q" .URL.Scheme }}
   req.Body = {{ .Req_Body }}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

var req_body = strings.NewReader({{ .Raw_Req_Body }})
`

type values struct {
   *http.Request
   Query url.Values
   Req_Body string
   Raw_Req_Body string
}

func Write_Request(req *http.Request, dst io.Writer) error {
   var v values
   if req.Body != nil && req.Method != "GET" {
      body, err := io.ReadAll(req.Body)
      if err != nil {
         return err
      }
      req.Body = io.NopCloser(bytes.NewReader(body))
      v.Raw_Req_Body = backquote(string(body))
      v.Req_Body = "io.NopCloser(req_body)"
   } else {
      v.Raw_Req_Body = `""`
      v.Req_Body = "nil"
   }
   v.Query = req.URL.Query()
   v.Request = req
   temp, err := new(template.Template).Parse(raw_temp)
   if err != nil {
      return err
   }
   return temp.Execute(dst, v)
}
func main() {
   var f flags
   flag.StringVar(&f.name, "f", "", "input file")
   flag.BoolVar(&f.golang, "g", false, "request as Go code")
   flag.StringVar(&f.output, "o", "", "output file")
   flag.BoolVar(&f.https, "s", false, "HTTPS")
   flag.Parse()
   if f.name != "" {
      create := os.Stdout
      if f.output != "" {
         var err error
         create, err = os.Create(f.output)
         if err != nil {
            panic(err)
         }
         defer create.Close()
      }
      open, err := os.Open(f.name)
      if err != nil {
         panic(err)
      }
      defer open.Close()
      req, err := read_request(open)
      if err != nil {
         panic(err)
      }
      if req.URL.Scheme == "" {
         if f.https {
            req.URL.Scheme = "https"
         } else {
            req.URL.Scheme = "http"
         }
      }
      if f.golang {
         err := Write_Request(req, create)
         if err != nil {
            panic(err)
         }
      } else {
         err := write(req, create)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
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

func backquote(s string) string {
   if strconv.Can_Backquote(s) {
      return "`" + s + "`"
   }
   return strconv.Quote(s)
}

type flags struct {
   golang bool
   https bool
   name string
   output string
}
