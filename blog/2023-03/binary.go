package binary

import (
   "encoding/hex"
   "unicode/utf8"
)

func encode(p []byte) []byte {
   var b []byte
   for len(p) >= 1 {
      if binary_byte(p[0]) {
         b = append(b, '=')
         b = append(b, hex.EncodeToString(p[:1])...)
         p = p[1:]
      } else {
         r, size := utf8.DecodeRune(p)
         src := p[:size]
         if r == utf8.RuneError && size == 1 {
            b = append(b, '=')
            b = append(b, hex.EncodeToString(src)...)
         } else {
            b = append(b, src...)
         }
         p = p[size:]
      }
   }
   return b
}

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
