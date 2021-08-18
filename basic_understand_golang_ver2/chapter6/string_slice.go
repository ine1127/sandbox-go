package main

import "fmt"

func main() {
  var x string = "abcde"[1:4]

  var y string = "あいうえお"[3:9]

  var z string = "あいうえお"[1:4]

  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
}
