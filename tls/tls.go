package tls

import (
   "crypto/md5"
   "encoding/binary"
   "encoding/hex"
   "github.com/refraction-networking/utls"
   "io"
   "net"
   "net/http"
)

func Fingerprint(ja3 string) string {
   sum := md5.Sum([]byte(ja3))
   return hex.EncodeToString(sum[:])
}

// this is currently the shortest one
const Android_API = Android_API_26

const Android_API_24 =
   "771,49195-49196-52393-49199-49200-52392-158-159-49161-49162-49171" +
   "-49172-51-57-156-157-47-53,65281-0-23-35-13-16-11-10,23,0"

const Android_API_25 =
   "771,49195-49196-52393-49199-49200-52392-158-159-49161-49162-49171" +
   "-49172-51-57-156-157-47-53,65281-0-23-35-13-16-11-10,23-24-25,0"
   
const Android_API_26 =
   "771,49195-49196-52393-49199-49200-52392-49161-49162-49171" +
   "-49172-156-157-47-53,65281-0-23-35-13-5-16-11-10,29-23-24,0"

const Android_API_29 =
   "771,4865-4866-4867-49195-49196-52393-49199-49200-52392-49161-49162-49171" +
   "-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-51-45-43-21,29-23-24,0"

const Android_API_32 = Android_API_29
func extension_type(ext tls.TLSExtension) (uint16, error) {
   pad, ok := ext.(*tls.UtlsPaddingExtension)
   if ok {
      pad.WillPad = true
      ext = pad
   }
   buf, err := io.ReadAll(ext)
   if err != nil || len(buf) <= 1 {
      return 0, err
   }
   return binary.BigEndian.Uint16(buf), nil
}

type Client_Hello struct {
   *tls.ClientHelloSpec
}

func New_Client_Hello() Client_Hello {
   var c Client_Hello
   c.ClientHelloSpec = new(tls.ClientHelloSpec)
   return c
}

// cannot call pointer method RoundTrip on http.Transport
func (c Client_Hello) Transport() *http.Transport {
   var tr http.Transport
   //lint:ignore SA1019 godocs.io/context
   tr.DialTLS = func(network, ref string) (net.Conn, error) {
      dial_conn, err := net.Dial(network, ref)
      if err != nil {
         return nil, err
      }
      host, _, err := net.SplitHostPort(ref)
      if err != nil {
         return nil, err
      }
      config := &tls.Config{ServerName: host}
      tls_conn := tls.UClient(dial_conn, config, tls.HelloCustom)
      if err := tls_conn.ApplyPreset(c.ClientHelloSpec); err != nil {
         return nil, err
      }
      if err := tls_conn.Handshake(); err != nil {
         return nil, err
      }
      return tls_conn, nil
   }
   return &tr
}
