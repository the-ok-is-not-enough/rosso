package strconv

import (
   "fmt"
   "testing"
   "unicode/utf8"
)

func Test_Valid(t *testing.T) {
   const s = "\xE0<"
   for _, r := range s {
      fmt.Println(utf8.ValidRune(r))
   }
   fmt.Println(utf8.ValidString(s))
}
