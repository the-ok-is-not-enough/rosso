package tls

import (
   "bytes"
   "crypto/md5"
   "encoding/hex"
)

func Fingerprint(ja3 []byte) []byte {
   hash := md5.New()
   hash.Write(ja3)
   src := hash.Sum(nil)
   dst := make([]byte, hex.EncodedLen(len(src)))
   hex.Encode(dst, src)
   return dst
}

// 8fcaa9e4a15f48af0a7d396e3fa5c5eb
func Android_API_24() []byte {
   var b bytes.Buffer
   b.WriteString("771,49195-49196-52393-49199-49200-52392-158-159")
   b.WriteString("-49161-49162-49171-49172-51-57-156-157-47-53")
   b.WriteString(",65281-0-23-35-13-16-11-10,23,0")
   return b.Bytes()
}

// 9fc6ef6efc99b933c5e2d8fcf4f68955
func Android_API_25() []byte {
   var b bytes.Buffer
   b.WriteString("771,49195-49196-52393-49199-49200-52392-158-159")
   b.WriteString("-49161-49162-49171-49172-51-57-156-157-47-53")
   b.WriteString(",65281-0-23-35-13-16-11-10,23-24-25,0")
   return b.Bytes()
}

// d8c87b9bfde38897979e41242626c2f3
func Android_API_26() []byte {
   var b bytes.Buffer
   b.WriteString("771,49195-49196-52393-49199-49200-52392")
   b.WriteString("-49161-49162-49171-49172-156-157-47-53")
   b.WriteString(",65281-0-23-35-13-5-16-11-10,29-23-24,0")
   return b.Bytes()
}

// 9b02ebd3a43b62d825e1ac605b621dc8
func Android_API_29() []byte {
   var b bytes.Buffer
   b.WriteString("771,4865-4866-4867-49195-49196-52393-49199-49200-52392")
   b.WriteString("-49161-49162-49171-49172-156-157-47-53")
   b.WriteString(",0-23-65281-10-11-35-16-5-13-51-45-43-21,29-23-24,0")
   return b.Bytes()
}

// same as API 29
func Android_API_32() []byte {
   return Android_API_29()
}

func Android_API() []byte {
   // this is currently the shortest one
   return Android_API_26()
}
