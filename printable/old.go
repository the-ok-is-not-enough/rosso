package printable

import "unicode/utf8"

func Binary(p []byte) bool {
   for _, b := range p {
      if Binary_Byte(b) {
         return true
      }
   }
   // ValidRune says each rune of "\xE0<" is valid, but as a string its not
   // valid UTF-8
   return !utf8.Valid(p)
}

// mimesniff.spec.whatwg.org#binary-data-byte
func Binary_Byte(b byte) bool {
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
