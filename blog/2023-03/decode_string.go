package printable

import (
   "encoding/hex"
   "strings"
   "unicode/utf8"
)

func decode_string(s string) ([]byte, error) {
   var b []byte
   for len(s) >= 1 {
      c := s[0]
      if c == escape_character {
         d, err := hex.DecodeString(s[1:3])
         if err != nil {
            return nil, err
         }
         b = append(b, d...)
         s = s[3:]
      } else {
         b = append(b, c)
         s = s[1:]
      }
   }
   return b, nil
}

func encode_string(p []byte) string {
   var b strings.Builder
   for len(p) >= 1 {
      r, size := decode_rune(p)
      src := p[:size]
      if r == utf8.RuneError && size == 1 {
         b.WriteByte(escape_character)
         hex.NewEncoder(&b).Write(src)
      } else {
         b.Write(src)
      }
      p = p[size:]
   }
   return b.String()
}
