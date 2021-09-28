package main

import (
  "flag"
  "fmt"
)

func main() {
  i := flag.Int("flag1", 123, "整数")
  s := flag.String("flag2", "abc", "文字列")

  flag.Parse()

  fmt.Println(*i, *s)
}
