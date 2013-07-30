package main

import( 
  "fmt"
  "regexp"
  "strings"
)

func find_index_of_first_vowel(s string) int{
  for i, c := range s{
    if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || (c == 'y' && i != 0) {
      return i
    }
  }
  return 0
}

func pig_latinize_sentence(s string) string {
  f := strings.Fields(s)
  m := make([]string, len(f))
  for i, v := range f {
    m[i] = pig_latinize_word( strings.ToLower(v) )
  }
  return strings.Join( m, " " )
}

func pig_latinize_word(w string) string {
  if regexp.MustCompile(`^[aeiou]`).MatchString(w){
    return w + "way"
  } else {
    var i int
    i = find_index_of_first_vowel(w)
    return w[i:] + w[:i] + "ay"  
  }
}

func main(){
  fmt.Println( pig_latinize_sentence( "The happy duck lays eight yellow eggs in a regular rhythem") )
}