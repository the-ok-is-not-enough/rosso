package main

import (
   "fmt"
   "text/template"
)

func main() {
   {
      s := template.JSEscapeString("\x00")
      fmt.Println(s)
   }
   {
      s := template.JSEscapeString("☺")
      fmt.Println(s)
   }
   {
      s := template.JSEscapeString("😀")
      fmt.Println(s)
   }
   {
      s := template.JSEscapeString("\n")
      fmt.Println(s)
   }
}
