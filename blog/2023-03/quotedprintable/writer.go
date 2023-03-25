package quotedprintable

const upperhex = "0123456789ABCDEF"

func isWhitespace(b byte) bool {
   return b == ' ' || b == '\t'
}

func encode(b byte) []byte {
   var p []byte
   p = append(p, '=')
   p = append(p, upperhex[b>>4])
   p = append(p, upperhex[b&0x0f])
   return p
}

func Write(p []byte) []byte {
   var q []byte
   for _, b := range p {
      switch {
      case b >= '!' && b <= '~' && b != '=':
         q = append(q, b)
      case isWhitespace(b) || (b == '\n' || b == '\r'):
         q = append(q, b)
      default:
         q = append(q, encode(b)...)
      }
   }
   return q
}

