# escape

we want to avoid this:

~~~go
package main

import "fmt"

func main() {
   b := []byte{1}
   fmt.Println(b) // [1]
   s := "\x01"
   fmt.Println(s) // ^A
}
~~~

search here:

https://cs.opensource.google/go/go/+/master:api

for:

~~~
\(string\).string
~~~

results:

- https://godocs.io/html#EscapeString
- https://godocs.io/html/template#HTMLEscapeString
- https://godocs.io/html/template#JSEscapeString
- https://godocs.io/mime/quotedprintable#Writer.Write
- https://godocs.io/net/url#PathEscape
- https://godocs.io/net/url#QueryEscape
- https://godocs.io/strconv#Quote
- https://godocs.io/strconv#QuoteToASCII
- https://godocs.io/strconv#QuoteToGraphic
- https://godocs.io/strings#Replacer.Replace
- https://godocs.io/text/template#HTMLEscapeString
- https://godocs.io/text/template#JSEscapeString

this looks promising:

https://godocs.io/mime/quotedprintable#Writer.Write
