package main

import "fmt"

func main() {
  var i interface{} = "test"

  var s string = i.(string)

  fmt.Println(s)
}
