package main

import (
   "2a.pages.dev/rosso/tls"
   "fmt"
   "net"
   "net/http"
   "os"
   utls "github.com/refraction-networking/utls"
)

func main() {
   var (
      hello tls.Client_Hello
      tr http.Transport
   )
   err := hello.UnmarshalText(tls.)
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
      tls_conn := tls.UClient(conn{Conn: dial_conn}, config, tls.HelloCustom)
      if err := tls_conn.ApplyPreset(hello); err != nil {
         return nil, err
      }
      if err := tls_conn.Handshake(); err != nil {
         return nil, err
      }
      return tls_conn, nil
   }
   req, err := http.NewRequest("GET", "https://mail.google.com", nil)
   if err != nil {
      panic(err)
   }
   res, err := tr.RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}

type conn struct {
   Handshake_Complete bool
   net.Conn
}

func (c conn) Write(b []byte) (int, error) {
   fmt.Printf("%q\n\n", b)
   return c.Conn.Write(b)
}
