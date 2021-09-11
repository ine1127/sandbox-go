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

  target := getElement(l, 3)

  fmt.Println("変更前:", target.Value)

  target.Value = "change"

  for e := l.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value)
  }
}

func getElement(l *list.List, index int) *list.Element {
  for e, i := l.Front(), 0; e != nil; e = e.Next() {
    if i == index {
      return e
    }

    i++
  }

  panic("インデックス不正")
}
