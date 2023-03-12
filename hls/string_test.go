package hls

import (
   "bytes"
   "fmt"
   "os"
   "sort"
   "testing"
)

func Test_Info(t *testing.T) {
   for test := range tests {
      buf, err := os.ReadFile(test + ".txt")
      if err != nil {
         t.Fatal(err)
      }
      sort.SliceStable(buf, func(int, int) bool {
         return true
      })
      master, err := New_Scanner(bytes.NewReader(buf)).Master()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(test)
      for _, item := range master.Streams {
         fmt.Println(item)
      }
      for _, item := range master.Media {
         fmt.Println(item)
      }
      fmt.Println()
   }
}
