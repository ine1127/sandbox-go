package main

import (
  "fmt"
  "hello/hello"
  "strconv"
)

func main() {
  x := hello.Input("type 1~5")
  n, err := strconv.Atoi(x)

  if err == nil {
    fmt.Print(x + "までの合計は、")
  } else {
    return
  }
  t := 0
  switch n {
  case 5:
    t += 5
    fallthrough
  case 4:
    t += 4
    fallthrough
  case 3:
    t += 3
    fallthrough
  case 2:
    t += 2
    fallthrough
  case 1:
    t++
  default:
    fmt.Println("範囲外です。")
    return
  }
  fmt.Println(t, "です。")
}
