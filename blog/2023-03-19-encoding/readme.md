# Encoding

this is attractive:

https://godocs.io/encoding#TextUnmarshaler

but in practice, text input is going to be a string, which means you will need
to cast to byte slice. Further, since its a method, you would need to initialize
the value:

~~~go
package main

import (
   "fmt"
   "time"
)

func main() {
   var t time.Time
   err := t.UnmarshalText([]byte("2006-01-02T15:04:05Z"))
   if err != nil {
      panic(err)
   }
   fmt.Println(t)
}
~~~

Implementing this interface does enable these functions:

- https://godocs.io/encoding/json#Unmarshal
- https://godocs.io/encoding/xml#Unmarshal
- https://godocs.io/flag#TextVar

but those are not useful to me at this time. Looks like typically a parse
function is written instead:

https://godocs.io/net#ParseIP

which is then wrapped with an `UnmarshalText` method if needed:

https://godocs.io/net#IP.UnmarshalText

also, a String method is written:

https://godocs.io/net#IP.String

which is then wrapped with a `MarshalText` method if needed:

https://godocs.io/net#IP.MarshalText
