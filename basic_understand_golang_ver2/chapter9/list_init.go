package main

import (
  "container/list"
  "fmt"
)

func main() {
  l := list.New()

  for i := 0; i < 5; i++ {
    l.PushBack(i)
  }

  l.Init()

  fmt.Println("要素数:", l.Len())
}
