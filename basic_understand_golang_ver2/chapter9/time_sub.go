package main

import (
  "fmt"
  "time"
)

func main() {
  start := time.Now()

  time.Sleep(time.Second * 15)

  end := time.Now()

  sub := end.Sub(start)

  fmt.Println(sub)
  fmt.Println(int(sub / time.Second), "秒")
}
