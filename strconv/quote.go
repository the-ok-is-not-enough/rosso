package strconv

import (
   "strconv"
   "unicode/utf8"
)

var (
   AppendQuote = strconv.AppendQuote
   Quote = strconv.Quote
)

func Valid(p []byte) bool {
   for _, item := range p {
      if binary(item) {
         return false
      }
   }
   // ValidRune says each rune of "\xE0<" is valid, but as a string its not
   // valid UTF-8
   return utf8.Valid(p)
}

// mimesniff.spec.whatwg.org#binary-data-byte
func binary[T byte|rune](item T) bool {
   if item <= 0x08 {
      return true
   }
   if item == 0x0B {
      return true
   }
   if item >= 0x0E && item <= 0x1A {
      return true
   }
   if item >= 0x1C && item <= 0x1F {
      return true
   }
   return false
}

