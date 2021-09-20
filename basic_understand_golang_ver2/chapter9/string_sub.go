package main

import (
  "fmt"
)

func main() {
  s := "abcdefghijklmn"

  fmt.Println(s[:10])
  fmt.Println(s[3:10])
  fmt.Println(s[3:])
  fmt.Println(s[:])
}
