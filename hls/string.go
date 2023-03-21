package hls

import (
   "strconv"
   "strings"
)

func (Medium) Ext() string {
   return ".m4a"
}

func (m Medium) URI() string {
   return m.Raw_URI
}

func (Stream) Ext() string {
   return ".m4v"
}

func (m Stream) URI() string {
   return m.Raw_URI
}

type Medium struct {
   Group_ID string
   Name string
   Raw_URI string
   Type string
   Characteristics string
}

type Stream struct {
   Bandwidth int64
   Raw_URI string
   Audio string
   Codecs string
   Resolution string
}

func (m Stream) String() string {
   var b []byte
   b = append(b, "bandwidth: "...)
   b = strconv.AppendInt(b, m.Bandwidth, 10)
   if m.Resolution != "" {
      b = append(b, "\n\tresolution: "...)
      b = append(b, m.Resolution...)
   }
   if m.Codecs != "" {
      b = append(b, "\n\tcodecs: "...)
      b = append(b, m.Codecs...)
   }
   if m.Audio != "" {
      b = append(b, "\n\taudio: "...)
      b = append(b, m.Audio...)
   }
   return string(b)
}

func (m Medium) String() string {
   var b strings.Builder
   b.WriteString("group ID: ")
   b.WriteString(m.Group_ID)
   b.WriteString("\n\ttype: ")
   b.WriteString(m.Type)
   b.WriteString("\n\tname: ")
   b.WriteString(m.Name)
   if m.Characteristics != "" {
      b.WriteString("\n\tcharacteristics: ")
      b.WriteString(m.Characteristics)
   }
   return b.String()
}
