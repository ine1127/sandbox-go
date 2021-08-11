package main

import "fmt"

func main() {
  var i interface{} = "test"

  s1, ok := i.(string)
  fmt.Println(s1, ok)

  s2, ok := i.(interface {
    dummy()
  })
  fmt.Println(s2, ok)
}
