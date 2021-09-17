package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()

  fmt.Println(t)

  add := t.AddDate(1, 0, 2)

  sub := t.AddDate(-1, 0, -2)

  fmt.Println(add)
  fmt.Println(sub)
}
