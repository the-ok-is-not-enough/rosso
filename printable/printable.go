package printable

import (
   "encoding/hex"
   "errors"
   "unicode/utf8"
   "strconv"
)

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

func Encode(src []byte) string {
   var dst []byte
   for len(src) >= 1 {
      r, size := decode_rune(src)
      s := src[:size]
      if r == utf8.RuneError && size == 1 {
         var d [2]byte
         hex.Encode(d[:], s)
         dst = append(dst, escape_character)
         dst = append(dst, d[:]...)
      } else {
         dst = append(dst, s...)
      }
      src = src[size:]
   }
   return string(dst)
}

const escape_character = '~'

var error_escape = errors.New("invalid printable escape")

func decode(src string) ([]byte, error) {
   var dst []byte
   for len(src) >= 1 {
      s := src[0]
      if s == escape_character {
         if len(src) <= 2 {
            return nil, error_escape
         }
         d, err := hex.DecodeString(src[1:3])
         if err != nil {
            return nil, err
         }
         dst = append(dst, d...)
         src = src[3:]
      } else {
         dst = append(dst, s)
         src = src[1:]
      }
   }
   return dst, nil
}

func decode_rune(p []byte) (rune, int) {
   if len(p) >= 1 {
      b := p[0]
      if b == escape_character || Binary_Byte(b) {
         return utf8.RuneError, 1
      }
   }
   return utf8.DecodeRune(p)
}
type unit_measure struct {
   factor float64
   name string
}

func (n Number) label(dst []byte, unit unit_measure) []byte {
   var prec int
   if unit.factor != 1 {
      prec = 2
   }
   unit.factor *= float64(n)
   dst = strconv.AppendFloat(dst, unit.factor, 'f', prec, 64)
   return append(dst, unit.name...)
}

func (n Number) scale(dst []byte, units []unit_measure) []byte {
   var unit unit_measure
   for _, unit = range units {
      if unit.factor * float64(n) < 1000 {
         break
      }
   }
   return n.label(dst, unit)
}

type Number float64

func New_Number[T Ordered](value T) Number {
   return Number(value)
}

func Ratio[T, U Ordered](num T, den U) Number {
   return Number(num) / Number(den)
}

type Ordered interface {
   ~float32 | ~float64 |
   ~int | ~int8 | ~int16 | ~int32 | ~int64 |
   ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func (n Number) Cardinal(dst []byte) []byte {
   units := []unit_measure{
      {1, ""},
      {1e-3, " thousand"},
      {1e-6, " million"},
      {1e-9, " billion"},
      {1e-12, " trillion"},
   }
   return n.scale(dst, units)
}

func (n Number) Percent(dst []byte) []byte {
   unit := unit_measure{100, "%"}
   return n.label(dst, unit)
}

func (n Number) Rate(dst []byte) []byte {
   units := []unit_measure{
      {1, " byte/s"},
      {1e-3, " kilobyte/s"},
      {1e-6, " megabyte/s"},
      {1e-9, " gigabyte/s"},
      {1e-12, " terabyte/s"},
   }
   return n.scale(dst, units)
}

func (n Number) Size(dst []byte) []byte {
   units := []unit_measure{
      {1, " byte"},
      {1e-3, " kilobyte"},
      {1e-6, " megabyte"},
      {1e-9, " gigabyte"},
      {1e-12, " terabyte"},
   }
   return n.scale(dst, units)
}
