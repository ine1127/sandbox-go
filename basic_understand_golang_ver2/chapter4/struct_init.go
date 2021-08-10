package main

import "fmt"

type Person struct {
  name string
  age int
}

func main() {
  var p1 Person
  p1.name = "Jhon"
  p1.age = 23

  p2 := Person{age: 31, name: "Tom"}

  p3 := Person{"Jane", 42}

  p4 := &Person{"Mike", 36}

  fmt.Println(p1, p2, p3, p4)
}
