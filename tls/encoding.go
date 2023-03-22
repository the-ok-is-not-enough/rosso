package tls

import (
   "github.com/refraction-networking/utls"
   "strconv"
   "strings"
)

func Parse(ja3 string) (*Client_Hello, error) {
   var (
      extensions string
      hello Client_Hello
      info tls.ClientHelloInfo
   )
   hello.ClientHelloSpec = new(tls.ClientHelloSpec)
   for i, field := range strings.SplitN(ja3, ",", 5) {
      switch i {
      case 0:
         // TLSVersMin is the record version, TLSVersMax is the handshake
         // version
         v, err := strconv.ParseUint(field, 10, 16)
         if err != nil {
            return nil, err
         }
         hello.TLSVersMax = uint16(v)
      case 1:
         // build CipherSuites
         for _, s := range strings.Split(field, "-") {
            v, err := strconv.ParseUint(s, 10, 16)
            if err != nil {
               return nil, err
            }
            hello.CipherSuites = append(hello.CipherSuites, uint16(v))
         }
      case 2:
         extensions = field
      case 3:
         for _, s := range strings.Split(field, "-") {
            v, err := strconv.ParseUint(s, 10, 16)
            if err != nil {
               return nil, err
            }
            info.SupportedCurves = append(info.SupportedCurves, tls.CurveID(v))
         }
      case 4:
         for _, s := range strings.Split(field, "-") {
            v, err := strconv.ParseUint(s, 10, 8)
            if err != nil {
               return nil, err
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
            return nil, err
         }
         ext = &tls.GenericExtension{Id: uint16(v)}
      }
      hello.Extensions = append(hello.Extensions, ext)
   }
   // uTLS does not support 0x0 as min version
   hello.TLSVersMin = tls.VersionTLS10
   return &hello, nil
}
