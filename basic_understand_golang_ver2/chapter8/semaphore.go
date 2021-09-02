package main

import (
  "fmt"
  "math/rand"
  "time"
)

const goroutines = 10

const maxProcesses = 3

func main() {
  semaphore := make(chan int, maxProcesses)

  notify := make(chan int, goroutines)

  for i := 0; i < goroutines; i++ {
    go func(no int, semaphore chan int, notify chan<- int) {
      semaphore <- 0
      time.Sleep(time.Duration(rand.Int63n(3)) * time.Second)
      fmt.Println("処理完了", no)

      <-semaphore
      notify <- no
    }(i, semaphore, notify)
  }

  for i := 0; i < goroutines; i++ {
    <-notify
  }

  fmt.Println("全て完了")
}
