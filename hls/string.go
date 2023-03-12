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

func (m Medium) String() string {
   var b strings.Builder
   b.WriteString("group ID: ")
   b.WriteString(m.Group_ID)
   b.WriteString("\nname: ")
   b.WriteString(m.Name)
   b.WriteString("\ntype: ")
   b.WriteString(m.Type)
   if m.Characteristics != "" {
      b.WriteString("\ncharacteristics: ")
      b.WriteString(m.Characteristics)
   }
   return b.String()
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
   if m.Audio != "" {
      b = append(b, "\naudio: "...)
      b = append(b, m.Audio...)
   }
   if m.Codecs != "" {
      b = append(b, "\ncodecs: "...)
      b = append(b, m.Codecs...)
   }
   if m.Resolution != "" {
      b = append(b, "\nresolution: "...)
      b = append(b, m.Resolution...)
   }
   return string(b)
}
