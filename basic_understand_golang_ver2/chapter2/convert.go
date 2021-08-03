package main

import "fmt"

func main() {
  var i int = 1234

  var u uint32 = uint32(i)

  var f float32 = float32(u)

  var s string = string(i)

  var b []byte = []byte("abc")

  fmt.Println(u)
  fmt.Println(f)
  fmt.Println(s)
  fmt.Println(b)
}
