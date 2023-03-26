package printable

import "unicode/utf8"

const escape_character = '~'

// mimesniff.spec.whatwg.org#binary-data-byte
func binary_byte(b byte) bool {
   if b <= 0x08 {
      return true
   }
   if b == 0x0B {
      return true
   }
   if b >= 0x0E && b <= 0x1A {
      return true
   }
   if b >= 0x1C && b <= 0x1F {
      return true
   }
   return false
}

func decode_rune(p []byte) (rune, int) {
   if len(p) >= 1 {
      b := p[0]
      if b == escape_character || binary_byte(b) {
         return utf8.RuneError, 1
      }
   }
   return utf8.DecodeRune(p)
}
