package main

import "fmt"

func main() {
  capitals := make(map[string]string)

  capitals["日本"] = "東京"
  capitals["アメリカ"] = "ワシントンD.C."
  capitals["中国"] = "北京"

  fmt.Println(capitals["日本"])
  fmt.Println(capitals["アメリカ"])
  fmt.Println(capitals["中国"])
  fmt.Println()

  for country, capital := range capitals {
    fmt.Println(country, capital)
  }
}
