package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Index("The Go Programming Lnaguage", "Go"))

  fmt.Println(strings.Index("Thr Go Programming Language", "xxx"))
}
