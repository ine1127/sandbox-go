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
    fmt.Print("1から" + x + "の偶数の合計は、")
  } else {
    return
  }
  t := 0
  c := 0

  for {
    c++
    if c % 2 == 1 {
      continue
    }
    if c > n {
      break
    }
    t += c
  }

  fmt.Println(t, "です。")
}
