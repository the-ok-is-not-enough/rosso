package main

import (
   "2a.pages.dev/rosso/http"
   "bufio"
   "flag"
   "os"
)

func main() {
   var f flags
   flag.StringVar(&f.name, "f", "", "input file")
   flag.BoolVar(&f.golang, "g", false, "request as Go code")
   flag.StringVar(&f.output, "o", "", "output file")
   flag.BoolVar(&f.https, "s", false, "HTTPS")
   flag.Parse()
   if f.name != "" {
      create := os.Stdout
      if f.output != "" {
         var err error
         create, err = os.Create(f.output)
         if err != nil {
            panic(err)
         }
         defer create.Close()
      }
      open, err := os.Open(f.name)
      if err != nil {
         panic(err)
      }
      defer open.Close()
      req, err := http.Read_Request(bufio.NewReader(open))
      if err != nil {
         panic(err)
      }
      if req.URL.Scheme == "" {
         if f.https {
            req.URL.Scheme = "https"
         } else {
            req.URL.Scheme = "http"
         }
      }
      if f.golang {
         err := write_request(req, create)
         if err != nil {
            panic(err)
         }
      } else {
         err := write(req, create)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
