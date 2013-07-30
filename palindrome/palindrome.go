package main

import(
  "fmt"
  "strings"
)

func is_palindrome(s string) bool{
  a := []rune(strings.ToLower(s))
  l := len(a) - 1 
  for i := 0; i <= l/2; i++ {
    if a[i] != a[l]{
      return false
    }
    l--
  }
  return true 
}

func main(){
  
  fmt.Printf( "Haninah: %t \n", is_palindrome("Haninah") )
  fmt.Printf( "Hannah: %t \n", is_palindrome("Hannah") )
  fmt.Printf( "Haninni: %t \n", is_palindrome("Haninni") )
  fmt.Printf( "世界: %t \n", is_palindrome("世界") )
  fmt.Printf( "世世: %t \n", is_palindrome("世世") )
  fmt.Printf( "世界世: %t \n", is_palindrome("世界世") )
  
}