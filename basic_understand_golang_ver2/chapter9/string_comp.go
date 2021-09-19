package main

import (
  "fmt"
)

func main() {
  s1 := "alpha"
  s2 := "beta"

  if s1 == s2 {
    fmt.Println("同じ")
  }

  if s1 != s2 {
    fmt.Println("異なる")
  }

  if s1 < s2 {
    fmt.Println("s1の方が小さい")
  }

  if s1 > s2 {
    fmt.Println("s2の方が小さい")
  }

  if s1 <= s2 {
    fmt.Println("s1の方が小さいか同じ")
  }

  if s1 >= s2 {
    fmt.Println("s2の方が小さいか同じ")
  }
}
