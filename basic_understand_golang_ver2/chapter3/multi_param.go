package main

import "fmt"

func main() {
  f2(f1())
}

func f1() (int, string, float32) {
  return 0, "xyz", 3.14
}

func f2(a int, b string, c interface{}) {
  fmt.Println(a, b, c)
}
