package main

import "fmt"

func main() {
  type myInteger int

  var i myInteger = 123

  i += 1

  fmt.Println(i)

  type myStruct struct {
    a int
    b int
  }
}
