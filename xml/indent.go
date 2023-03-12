package xml

import (
   "bytes"
   "encoding/xml"
   "io"
)

func Indent(dst io.Writer, src io.Reader, prefix, indent string) error {
   decode := xml.NewDecoder(src)
   encode := xml.NewEncoder(dst)
   encode.Indent(prefix, indent)
   for {
      token, err := decode.Token()
      if err == io.EOF {
         return encode.Flush()
      }
      if err != nil {
         return err
      }
      data, ok := token.(xml.CharData)
      if ok {
         token = xml.CharData(bytes.TrimSpace(data))
      }
      if err := encode.EncodeToken(token); err != nil {
         return err
      }
   }
}
