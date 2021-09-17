package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()

  switch t.Weekday() {
  case time.Sunday:
    fmt.Println("日曜日")
  case time.Monday:
    fmt.Println("月曜日")
  case time.Tuesday:
    fmt.Println("火曜日")
  case time.Wednesday:
    fmt.Println("水曜日")
  case time.Thursday:
    fmt.Println("木曜日")
  case time.Friday:
    fmt.Println("金曜日")
  case time.Saturday:
    fmt.Println("土曜日")
  }
}
