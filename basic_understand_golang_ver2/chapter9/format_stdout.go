package main

import "fmt"

func main() {
  fmt.Print("東京", "の降水確率は", 90, "%です。\n")
  fmt.Println("東京", "の降水確率は", 90, "%です。")
  fmt.Printf("%sの降水確率は%d%%です。\n", "東京", 90)
}
