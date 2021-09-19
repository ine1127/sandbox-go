package main

import (
  "fmt"
  "strings"
)

func main() {
  s := "alpha"
  s = s + "beta"
  s += "gamma"

  fmt.Println(s)

  arr := []string{"alpha", "beta", "gamma"}

  fmt.Println(strings.Join(arr, ","))
}
