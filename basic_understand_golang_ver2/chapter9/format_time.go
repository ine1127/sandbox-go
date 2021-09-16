package main

import (
  "fmt"
  "time"
)

func main() {
  time := time.Now()

  fmt.Printf("%04d/%02d/%02d %02d:%02d:%02d\n",
    time.Year(),
    time.Month(),
    time.Day(),
    time.Hour(),
    time.Minute(),
    time.Second(),
  )
}
