package main

import (
  "fmt"
  "hello/hello"
  "strconv"
)

func main() {
  x := hello.Input("type a number")
  n, err := strconv.Atoi(x)

  if err == nil {
    fmt.Print("1から" + x + "の合計は、")
  } else {
    return
  }
  t := 0
  c := 1

  for c <= n {
    t += c
    c++
  }

  fmt.Println(t, "です。")
}
