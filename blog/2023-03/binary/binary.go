package quotedprintable

func escape(p []byte) []byte {
   var s []byte
   for _, b := range p {
      if b == '=' {
         s = append(s, "=3D"...)
      } else if binary_byte(b) {
         s = append(s, '=')
         s = append(s, upperhex[b>>4])
         s = append(s, upperhex[b&0x0f])
      } else {
         s = append(s, b)
      }
   }
   return s
}

const upperhex = "0123456789ABCDEF"

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
