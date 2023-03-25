package main

import (
   "encoding/xml"
   "fmt"
   "html"
   "mime/quotedprintable"
   "net/url"
   "strconv"
   "strings"
   "text/template"
)

const s = "\x01¶'"

var escapes = []string{
   // github.com/golang/go/blob/go1.20.2/src/encoding/xml/xml.go#L1902
   func() string {
      var b strings.Builder
      xml.EscapeText(&b, []byte(s))
      return b.String()
   }(),
   // github.com/golang/go/blob/go1.20.2/src/html/escape.go#L178
   html.EscapeString(s),
   // github.com/golang/go/blob/go1.20.2/src/mime/quotedprintable/writer.go#L31
   // github.com/golang/go/blob/go1.20.2/src/mime/quotedprintable/writer.go#L112
   func() string {
      var b strings.Builder
      w := quotedprintable.NewWriter(&b)
      w.Write([]byte(s))
      w.Close()
      return b.String()
   }(),
   // github.com/golang/go/blob/go1.20.2/src/net/url/url.go#L281
   url.PathEscape(s),
   // github.com/golang/go/blob/go1.20.2/src/net/url/url.go#L275
   url.QueryEscape(s),
   // github.com/golang/go/blob/go1.20.2/src/strconv/quote.go#L128
   strconv.Quote(s),
   // github.com/golang/go/blob/go1.20.2/src/strconv/quote.go#L141
   strconv.QuoteToASCII(s),
   // github.com/golang/go/blob/go1.20.2/src/strconv/quote.go#L155
   strconv.QuoteToGraphic(s),
   // github.com/golang/go/blob/go1.20.2/src/text/template/funcs.go#L611
   func() string {
      var b strings.Builder
      template.HTMLEscape(&b, []byte(s))
      return b.String()
   }(),
   // github.com/golang/go/blob/go1.20.2/src/text/template/funcs.go#L671
   func() string {
      var b strings.Builder
      template.JSEscape(&b, []byte(s))
      return b.String()
   }(),
}

func main() {
   for _, escape := range escapes {
      fmt.Println(escape)
   }
}
