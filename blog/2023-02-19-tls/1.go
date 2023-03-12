package main

import (
   "bufio"
   "crypto/tls"
   "fmt"
   "net"
   "net/http"
   "os"
)

type conn struct {
   net.Conn
}

func (c conn) Write(b []byte) (int, error) {
   fmt.Printf("%q\n", b)
   return c.Conn.Write(b)
}

func main() {
   req, err := http.NewRequest("GET", "https://mail.google.com", nil)
   if err != nil {
      panic(err)
   }
   dial_conn, err := net.Dial("tcp", "mail.google.com:443")
   if err != nil {
      panic(err)
   }
   config := tls.Config{ServerName: "mail.google.com"}
   tls_conn := tls.Client(conn{dial_conn}, &config)
   defer tls_conn.Close()
   if err := req.Write(tls_conn); err != nil {
      panic(err)
   }
   res, err := http.ReadResponse(bufio.NewReader(tls_conn), nil)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
