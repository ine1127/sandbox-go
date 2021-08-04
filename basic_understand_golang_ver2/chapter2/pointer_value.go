package main

import "fmt"

func main() {
  a, b := 1, 1

  double(a, &b)

  fmt.Println("値渡し:", a)
  fmt.Println("ポインタ渡し:", b)
}

func double(x int, y *int) {
  x = x * 2
  *y = *y * 2
}
