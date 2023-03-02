package tls

import (
   "github.com/refraction-networking/utls"
   "strconv"
)

// encoding.TextMarshaler using JA3
func (c Client_Hello) MarshalText() ([]byte, error) {
   var b []byte
   // TLSVersMin is the record version, TLSVersMax is the handshake version
   b = strconv.AppendUint(b, uint64(c.TLSVersMax), 10)
   // Cipher Suites
   b = append(b, ',')
   for key, val := range c.CipherSuites {
      if key >= 1 {
         b = append(b, '-')
      }
      b = strconv.AppendUint(b, uint64(val), 10)
   }
   // Extensions
   b = append(b, ',')
   var (
      curves []tls.CurveID
      points []uint8
   )
   for key, val := range c.Extensions {
      switch ext := val.(type) {
      case *tls.SupportedCurvesExtension:
         curves = ext.Curves
      case *tls.SupportedPointsExtension:
         points = ext.SupportedPoints
      }
      typ, err := extension_type(val)
      if err != nil {
         return nil, err
      }
      if key >= 1 {
         b = append(b, '-')
      }
      b = strconv.AppendUint(b, uint64(typ), 10)
   }
   // Elliptic curves
   b = append(b, ',')
   for key, val := range curves {
      if key >= 1 {
         b = append(b, '-')
      }
      b = strconv.AppendUint(b, uint64(val), 10)
   }
   // ECPF
   b = append(b, ',')
   for key, val := range points {
      if key >= 1 {
         b = append(b, '-')
      }
      b = strconv.AppendUint(b, uint64(val), 10)
   }
   return b, nil
}
