package strconv

import "unicode/utf8"

// mimesniff.spec.whatwg.org#binary-data-byte
func valid_rune(r rune) bool {
   if r >= 0x00 && r <= 0x08 {
      return false
   }
   if r == 0x0B {
      return false
   }
   if r >= 0x0E && r <= 0x1A {
      return false
   }
   if r >= 0x1C && r <= 0x1F {
      return false
   }
   if r == 0xFFFD {
      return false
   }
   return utf8.ValidRune(r)
}
