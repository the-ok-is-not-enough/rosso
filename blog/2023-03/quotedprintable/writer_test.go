package quotedprintable

import (
   "fmt"
   "strings"
   "testing"
)

const s = "\x01¶'"

func Test_Write(t *testing.T) {
   var b strings.Builder
   w := NewWriter(&b)
   w.Write([]byte(s))
   w.Close()
   fmt.Println(b.String())
}
