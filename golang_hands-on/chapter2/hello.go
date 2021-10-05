package main

import (
  "fmt"
  "hello/hello"
  "strconv"
  "strings"
)

func main() {
  x := hello.Input("Input data")
  ar := strings.Split(x, " ")
  t := 0

  for i := 0; i < len(ar); i++ {
    n, er := strconv.Atoi(ar[i])

    if er != nil {
      goto err
    }
    t += n
  }

  fmt.Println("total:", t)
  return

err:
  fmt.Println("ERROR!")
}
