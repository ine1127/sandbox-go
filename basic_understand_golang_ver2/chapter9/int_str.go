package main

import (
  "fmt"
  "strconv"
)

func main() {
  var i64 int64 = 1234567890

  fmt.Println(strconv.FormatInt(i64, 10))
  fmt.Println(strconv.FormatInt(i64, 16))

  var i32 int32 = 1234567890

  fmt.Println(strconv.FormatInt(int64(i32), 10))
}
