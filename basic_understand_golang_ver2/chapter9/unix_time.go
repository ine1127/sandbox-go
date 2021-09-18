package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()

  unix := t.Unix()

  fmt.Println(unix)
}
