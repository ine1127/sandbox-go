package main

import (
  "fmt"
  "strings"
)

func main() {
  s := "one,two,three,four,five"

  result := strings.Split(s, ",")

  fmt.Printf("%q\n", result)
}
