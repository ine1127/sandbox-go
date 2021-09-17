package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()

  fmt.Println(t)

  add := t.Add(time.Hour*24 + time.Minute*30 + time.Second*5)

  sub := t.Add(-time.Hour*24 - time.Minute*30 - time.Second*5)

  fmt.Println(add)
  fmt.Println(sub)
}
