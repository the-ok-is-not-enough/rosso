package main

import (
   "the-ok-is-not-enough/rosso/protobuf"
   "encoding/json"
   "flag"
   "os"
)

type flags struct {
   input string
   output string
}

func main() {
   var f flags
   flag.StringVar(&f.input, "f", "", "input file")
   flag.StringVar(&f.output, "o", "", "output file")
   flag.Parse()
   
   if f.input != "" {
      err := f.protobuf()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func (f flags) protobuf() error {
   data, err := os.ReadFile(f.input)
   if err != nil {
      return err
   }
   mes, err := protobuf.Unmarshal(data)
   if err != nil {
      return err
   }
   file := os.Stdout
   if f.output != "" {
      file, err := os.Create(f.output)
      if err != nil {
         return err
      }
      defer file.Close()
   }
   enc := json.NewEncoder(file)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   return enc.Encode(mes)
}
