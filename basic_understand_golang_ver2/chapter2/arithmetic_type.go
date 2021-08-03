package main

import "fmt"

const C1 = 12345

const C2 int = 34567

func main() {
  var x int = C1 * C2

  fmt.Println("12345 * 34567", x)

  var a int32 = 123
  var b int64 = 234

  fmt.Println("123 + 234 =", a+int32(b))
}
