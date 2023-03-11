package hls

import (
   "fmt"
   "os"
   "testing"
)

func Test_Info(t *testing.T) {
   for key, val := range tests {
      file, err := os.Open(key)
      if err != nil {
         t.Fatal(err)
      }
      master, err := New_Scanner(file).Master()
      if err != nil {
         t.Fatal(err)
      }
      if err := file.Close(); err != nil {
         t.Fatal(err)
      }
      fmt.Println(key)
      for _, item := range master.Streams.Filter(val.stream) {
         fmt.Println(item)
      }
      for _, item := range master.Media.Filter(val.medium) {
         fmt.Println(item)
      }
      fmt.Println()
   }
}
