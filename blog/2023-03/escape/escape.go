package main

import (
   "encoding/xml"
   "fmt"
   "html"
   "html/template"
   "mime/quotedprintable"
   "net/url"
   "strconv"
   "strings"
   text "text/template"
)

const s = "\x01¶'"

func main() {
   var escapes = []string{
      0: html.EscapeString(s),
      1: strconv.Quote(s),
      2: strconv.QuoteToASCII(s),
      3: strconv.QuoteToGraphic(s),
      4: strings.NewReplacer("\x01", `\x01`).Replace(s),
      5: template.HTMLEscapeString(s),
      6: template.JSEscapeString(s),
      7: text.HTMLEscapeString(s),
      8: text.JSEscapeString(s),
      9: url.PathEscape(s),
      10: url.QueryEscape(s),
      11: func() string {
         var b strings.Builder
         xml.Escape(&b, []byte(s))
         return b.String()
      }(),
      12: func() string {
         var b strings.Builder
         w := quotedprintable.NewWriter(&b)
         w.Write([]byte(s))
         w.Close()
         return b.String()
      }(),
   }
   for i, escape := range escapes {
      fmt.Println(i, escape)
   }
}
