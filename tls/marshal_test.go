package tls

import (
   "bytes"
   "testing"
)

func Test_MarshalText(t *testing.T) {
   a := Android_API()
   hello := New_Client_Hello()
   err := hello.UnmarshalText(a)
   if err != nil {
      t.Fatal(err)
   }
   b, err := hello.MarshalText()
   if err != nil {
      t.Fatal(err)
   }
   if !bytes.Equal(b, a) {
      t.Fatal(b)
   }
}
