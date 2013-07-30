package main 

import(
  "fmt"
  "regexp"
  "unicode/utf8"
)

func count_vowels(s string) map[string]int {
  m := make( map[string]int, utf8.RuneCountInString(s) )
  vowels := []rune { 'a','e','i','o', 'u'}
  for _, c := range s {
    for _, v := range vowels {
      if v == c {
        m[string(c)]++
      }
    }
  }
  return m
}

func count_vowels_regex(s string) map[string]int {
  m     := make(map[string]int) 
  re, _ := regexp.Compile(`[aeiou]`)
  res   := re.FindAllString(s, -1)
  for _, v := range res {
    m[v]++ 
  }
  return m
}

func main() {
  fmt.Println( count_vowels("hellooooooooo") )
  fmt.Println( count_vowels_regex("hellooooooooo") )
}