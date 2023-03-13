package main

import (
   "2a.pages.dev/rosso/xml"
   "encoding/json"
   "flag"
   "os"
)

func (f flags) indent_xml() error {
   src, err := os.Open(f.input)
   if err != nil {
      return err
   }
   defer src.Close()
   dst := os.Stdout
   if f.output != "" {
      dst, err = os.Create(f.output)
      if err != nil {
         return err
      }
      defer dst.Close()
   }
   return xml.Indent(dst, src, "", " ")
}

func (f flags) indent_json() error {
   in, err := os.Open(f.input)
   if err != nil {
      return err
   }
   defer in.Close()
   out, err := os.Create(f.output)
   if err != nil {
      out = os.Stdout
   }
   defer out.Close()
   var value any
   if err := json.NewDecoder(in).Decode(&value); err != nil {
      return err
   }
   enc := json.NewEncoder(out)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   return enc.Encode(value)
}

type flags struct {
   input string
   output string
   xml bool
}

func main() {
   var f flags
   flag.StringVar(&f.input, "f", "", "input file")
   flag.StringVar(&f.output, "o", "", "output file")
   flag.BoolVar(&f.xml, "xml", false, "use XML instead of JSON")
   flag.Parse()
   if f.input != "" {
      var err error
      if f.xml {
         err = f.indent_xml()
      } else {
         err = f.indent_json()
      }
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

