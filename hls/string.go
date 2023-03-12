package hls

import (
   "bytes"
   "strconv"
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
   b := m.Marshal_Indent("\t")
   return string(b)
}

func (m Medium) String() string {
   b := m.Marshal_Indent("\t")
   return string(b)
}

//////////////////////////////////////////////////////////

func (m Stream) Marshal_Indent(indent string) []byte {
   var b []byte
   b = append(b, "bandwidth: "...)
   b = strconv.AppendInt(b, m.Bandwidth, 10)
   if m.Resolution != "" {
      b = append(b, '\n')
      b = append(b, indent...)
      b = append(b, "resolution: "...)
      b = append(b, m.Resolution...)
   }
   if m.Codecs != "" {
      b = append(b, '\n')
      b = append(b, indent...)
      b = append(b, "codecs: "...)
      b = append(b, m.Codecs...)
   }
   if m.Audio != "" {
      b = append(b, '\n')
      b = append(b, indent...)
      b = append(b, "audio: "...)
      b = append(b, m.Audio...)
   }
   return b
}

func (m Medium) Marshal_Indent(indent string) []byte {
   var b bytes.Buffer
   b.WriteString("group ID: ")
   b.WriteString(m.Group_ID)
   b.WriteByte('\n')
   b.WriteString(indent)
   b.WriteString("type: ")
   b.WriteString(m.Type)
   b.WriteByte('\n')
   b.WriteString(indent)
   b.WriteString("name: ")
   b.WriteString(m.Name)
   if m.Characteristics != "" {
      b.WriteByte('\n')
      b.WriteString(indent)
      b.WriteString("characteristics: ")
      b.WriteString(m.Characteristics)
   }
   return b.Bytes()
}
