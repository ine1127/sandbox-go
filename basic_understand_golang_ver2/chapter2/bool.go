package main

import "fmt"

func main() {
  var b bool

  b = true
  b = false

  // bool型の変数に論理演算した結果を代入
  b = true || false

  fmt.Println(b)
}
