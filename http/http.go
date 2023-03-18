package http

import (
   "2a.pages.dev/rosso/strconv"
   "bufio"
   "bytes"
   "errors"
   "io"
   "net/http"
   "net/http/httputil"
   "net/textproto"
   "net/url"
   "os"
   "strings"
   "time"
)

func (c Client) Do(req *http.Request) (*http.Response, error) {
   switch c.Log_Level {
   case 1:
      os.Stderr.WriteString(req.Method)
      os.Stderr.WriteString(" ")
      os.Stderr.WriteString(req.URL.String())
      os.Stderr.WriteString("\n")
   case 2:
      dump, err := httputil.DumpRequest(req, true)
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
   res, err := c.client.Do(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != c.status {
      return nil, errors.New(res.Status)
   }
   return res, nil
}

type Client struct {
   Log_Level int // this needs to work with flag.IntVar
   status int
   client http.Client
}

var Default_Client = Client{
   Log_Level: 1,
   client: http.Client{
      CheckRedirect: func(*http.Request, []*http.Request) error {
         return http.ErrUseLastResponse
      },
   },
   status: http.StatusOK,
}

func (c Client) Get(ref string) (*http.Response, error) {
   req, err := http.NewRequest("GET", ref, nil)
   if err != nil {
      return nil, err
   }
   return c.Do(req)
}

func (c Client) Level(level int) Client {
   c.Log_Level = level
   return c
}

func (c Client) Redirect(fn Redirect_Func) Client {
   c.client.CheckRedirect = nil
   return c
}

func (c Client) Status(status int) Client {
   c.status = status
   return c
}

func (c Client) Transport(tr *http.Transport) Client {
   c.client.Transport = tr
   return c
}

type Redirect_Func func(*http.Request, []*http.Request) error
var NewRequest = http.NewRequest

func Read_Request(in io.Reader) (*http.Request, error) {
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

type Cookie = http.Cookie

type Header = http.Header

type Progress struct {
   bytes int64
   bytes_read int64
   bytes_written int
   chunks int
   chunks_read int64
   lap time.Time
   total time.Time
   w io.Writer
}

func Progress_Bytes(dst io.Writer, bytes int64) *Progress {
   return &Progress{w: dst, bytes: bytes}
}

func Progress_Chunks(dst io.Writer, chunks int) *Progress {
   return &Progress{w: dst, chunks: chunks}
}

func (p *Progress) Add_Chunk(bytes int64) {
   p.bytes_read += bytes
   p.chunks_read += 1
   p.bytes = int64(p.chunks) * p.bytes_read / p.chunks_read
}

func (p *Progress) Write(data []byte) (int, error) {
   if p.total.IsZero() {
      p.total = time.Now()
      p.lap = time.Now()
   }
   lap := time.Since(p.lap)
   if lap >= time.Second {
      total := time.Since(p.total).Seconds()
      var b []byte
      b = strconv.Ratio(p.bytes_written, p.bytes).Percent(b)
      b = append(b, "   "...)
      b = strconv.New_Number(p.bytes_written).Size(b)
      b = append(b, "   "...)
      b = strconv.Ratio(p.bytes_written, total).Rate(b)
      b = append(b, '\n')
      os.Stderr.Write(b)
      p.lap = p.lap.Add(lap)
   }
   write, err := p.w.Write(data)
   p.bytes_written += write
   return write, err
}

type Request = http.Request

type Response = http.Response

type Transport = http.Transport
