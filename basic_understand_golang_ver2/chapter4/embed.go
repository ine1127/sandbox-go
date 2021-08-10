package main

import "fmt"

type embedded struct {
  i int
}

func (x embedded) doSomething() {
  fmt.Println("test.doSomething()")
}

type test struct {
  embedded
}

func main() {
  var x test

  x.doSomething()

  fmt.Println(x.i)
}
