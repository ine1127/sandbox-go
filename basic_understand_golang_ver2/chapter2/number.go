package main

import "fmt"

func main() {
  var i int = 12345

  // intからint64への変換
  var i64 int64 = int64(i)

  fmt.Println(i64)
}
