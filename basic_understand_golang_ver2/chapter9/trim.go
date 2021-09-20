package main

import (
  "fmt"
  "strings"
)

func main() {
  s := " xyz "

  fmt.Printf("[%s]\n", strings.TrimSpace(s))

  fmt.Printf("[%s]\n", strings.TrimLeft(s, " "))
  fmt.Printf("[%s]\n", strings.TrimRight(s, " "))
}
