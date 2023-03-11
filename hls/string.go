package hls

import (
   "strconv"
   "strings"
)

func (m Stream) String() string {
   var b []byte
   if m.Resolution != "" {
      b = append(b, "Resolution:"...)
      b = append(b, m.Resolution...)
      b = append(b, ' ')
   }
   b = append(b, "Bandwidth:"...)
   b = strconv.AppendInt(b, m.Bandwidth, 10)
   if m.Codecs != "" {
      b = append(b, " Codecs:"...)
      b = append(b, m.Codecs...)
   }
   if m.Audio != "" {
      b = append(b, "\n  Audio:"...)
      b = append(b, m.Audio...)
   }
   return string(b)
}

type Stream struct {
   Audio string
   Bandwidth int64
   Codecs string
   Resolution string
   Raw_URI string
}

func (Medium) Ext() string {
   return ".m4a"
}

func (m Medium) String() string {
   var buf strings.Builder
   buf.WriteString("Type:")
   buf.WriteString(m.Type)
   buf.WriteString(" Name:")
   buf.WriteString(m.Name)
   buf.WriteString("\n  Group ID:")
   buf.WriteString(m.Group_ID)
   if m.Characteristics != "" {
      buf.WriteString("\n  Characteristics:")
      buf.WriteString(m.Characteristics)
   }
   return buf.String()
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
   Characteristics string
   Group_ID string
   Name string
   Raw_URI string
   Type string
}
