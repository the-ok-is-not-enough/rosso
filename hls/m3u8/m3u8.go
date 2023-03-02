package main

import (
   "os"
   "sort"
)

var names = []string{
   "cbc-audio.m3u8",
   "cbc-master.m3u8",
   "cbc-video.m3u8",
   "nbc-master.m3u8",
   "nbc-segment.m3u8",
   "paramount-master.m3u8",
   "paramount-segment.m3u8",
   "roku-master.m3u8",
   "roku-segment.m3u8",
}

func main() {
   for _, name := range names {
      buf, err := os.ReadFile(name)
      if err != nil {
         panic(err)
      }
      sort.SliceStable(buf, func(int, int) bool {
         return true
      })
      if err := os.WriteFile(name + ".txt", buf, os.ModePerm); err != nil {
         panic(err)
      }
   }
}
