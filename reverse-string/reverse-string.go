package main 

import(
  "fmt"
  "unicode/utf8"
)

func reverse_string(s string) string {
  a := make([]rune, utf8.RuneCountInString(s))
  i := len(a) - 1 
  for _, c := range s {
    a[i] = c
    i--
  }
  return string(a)
}

func main() {
 fmt.Println( reverse_string("Hello World or 世界") ) 
}