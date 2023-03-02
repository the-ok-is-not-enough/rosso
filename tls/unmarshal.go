package tls

import (
   "encoding/binary"
   "github.com/refraction-networking/utls"
   "strconv"
   "strings"
)

// encoding.BinaryUnmarshaler
func (c Client_Hello) UnmarshalBinary(data []byte) error {
   // unsupported extension 0x16
   printer := tls.Fingerprinter{AllowBluntMimicry: true}
   var err error
   c.ClientHelloSpec, err = printer.FingerprintClientHello(data)
   if err != nil {
      return err
   }
   // If SupportedVersionsExtension is present, then TLSVersMax is set to zero.
   // In which case we need to manually read the bytes.
   if c.TLSVersMax == 0 {
      // \x16\x03\x01\x00\xbc\x01\x00\x00\xb8\x03\x03
      c.TLSVersMax = binary.BigEndian.Uint16(data[9:])
   }
   return nil
}

// encoding.TextUnmarshaler using JA3
func (c Client_Hello) UnmarshalText(text []byte) error {
   var (
      extensions string
      info tls.ClientHelloInfo
   )
   for i, field := range strings.SplitN(string(text), ",", 5) {
      switch i {
      case 0:
         // TLSVersMin is the record version, TLSVersMax is the handshake
         // version
         v, err := strconv.ParseUint(field, 10, 16)
         if err != nil {
            return err
         }
         c.TLSVersMax = uint16(v)
      case 1:
         // build CipherSuites
         for _, s := range strings.Split(field, "-") {
            v, err := strconv.ParseUint(s, 10, 16)
            if err != nil {
               return err
            }
            c.CipherSuites = append(c.CipherSuites, uint16(v))
         }
      case 2:
         extensions = field
      case 3:
         for _, s := range strings.Split(field, "-") {
            v, err := strconv.ParseUint(s, 10, 16)
            if err != nil {
               return err
            }
            info.SupportedCurves = append(info.SupportedCurves, tls.CurveID(v))
         }
      case 4:
         for _, s := range strings.Split(field, "-") {
            v, err := strconv.ParseUint(s, 10, 8)
            if err != nil {
               return err
            }
            info.SupportedPoints = append(info.SupportedPoints, uint8(v))
         }
      }
   }
   // build extenions list
   for _, s := range strings.Split(extensions, "-") {
      var ext tls.TLSExtension
      switch s {
      case "0":
         // Android API 24
         ext = &tls.SNIExtension{}
      case "5":
         // Android API 26
         ext = &tls.StatusRequestExtension{}
      case "10":
         ext = &tls.SupportedCurvesExtension{Curves: info.SupportedCurves}
      case "11":
         ext = &tls.SupportedPointsExtension{
            SupportedPoints: info.SupportedPoints,
         }
      case "13":
         ext = &tls.SignatureAlgorithmsExtension{
            SupportedSignatureAlgorithms: []tls.SignatureScheme{
               // Android API 24
               tls.ECDSAWithP256AndSHA256,
               // httpbin.org
               tls.PKCS1WithSHA256,
            },
         }
      case "16":
         // Android API 24
         ext = &tls.ALPNExtension{
            AlpnProtocols: []string{"http/1.1"},
         }
      case "23":
         // Android API 24
         ext = &tls.UtlsExtendedMasterSecretExtension{}
      case "27":
         // Google Chrome
         ext = &tls.UtlsCompressCertExtension{
            Algorithms: []tls.CertCompressionAlgo{tls.CertCompressionBrotli},
         }
      case "43":
         // Android API 29
         ext = &tls.SupportedVersionsExtension{
            Versions: []uint16{tls.VersionTLS12},
         }
      case "45":
         // Android API 29
         ext = &tls.PSKKeyExchangeModesExtension{
            Modes: []uint8{tls.PskModeDHE},
         }
      case "65281":
         // Android API 24
         ext = &tls.RenegotiationInfoExtension{}
      default:
         v, err := strconv.ParseUint(s, 10, 16)
         if err != nil {
            return err
         }
         ext = &tls.GenericExtension{Id: uint16(v)}
      }
      c.Extensions = append(c.Extensions, ext)
   }
   // uTLS does not support 0x0 as min version
   c.TLSVersMin = tls.VersionTLS10
   return nil
}
