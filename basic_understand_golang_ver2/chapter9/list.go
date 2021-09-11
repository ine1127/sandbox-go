package main

import (
  "container/list"
  "fmt"
)

func main() {
  l := list.New()

  l.PushBack("a")
  l.PushBack("b")
  l.PushBack(3)

  for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
  }
}
