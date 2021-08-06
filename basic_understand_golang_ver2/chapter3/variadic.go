package main

import "fmt"

func main() {
  holiday(1, "元旦", "成人の日")
  holiday(2, "建国記念の日")
  holiday(3, "春分の日")
}

func holiday(month int, days ...string) {
  fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))
  for _, day := range days {
    fmt.Println(day)
  }
  fmt.Println()
}
