package main

import "fmt"

func main() {
  s1 := []int{1, 2, 3, 4}

  s2 := append(s1, 5, 6)

  s3 := append(s2, s1...)

  fmt.Println(s3)

  var b1 []byte
  b2 := append(b1, "abc"...)
  fmt.Println(b2)
}
