package main

import "fmt"
import "unicode/utf8"

func main() {
  var ja string = "Go言語"

  fmt.Println(ja, "len:", utf8.RuneCountInString(ja))
}
