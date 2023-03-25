package quotedprintable

import (
   "fmt"
   "testing"
)

const s = "\x01¶\n'"

func Test_Write(t *testing.T) {
   p := Write([]byte(s))
   fmt.Println(string(p))
}
