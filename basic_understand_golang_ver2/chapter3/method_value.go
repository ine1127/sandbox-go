package main

import "fmt"

type myType int

// func (レシーバの変数 レシーバの型) メソッド名(パラメータリスト) 戻り値の型 {
func (value *myType) add(increment myType) myType {
  *value += increment
  return *value
}

func main() {
  var i myType

  fmt.Println(i.add(3))

  mv := i.add

  fmt.Println(mv(3))
}
