package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("main:開始")
  fmt.Println("testを通常の関数として起動")
  test()

  fmt.Println("testをゴルーチンとして起動")
  go test()

  time.Sleep(3 * time.Second)

  fmt.Println("main:終了")
}

func test() {
  for i := 0; i < 5; i++ {
    fmt.Println(i)
    time.Sleep(1 * time.Second)
  }
}
