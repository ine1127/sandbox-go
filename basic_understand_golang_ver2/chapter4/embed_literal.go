package main

import "fmt"

type Person struct {
  name string
  age int
}

type Employee struct {
  id int
  Person
}

func main() {
  e := Employee{1, Person{"Jack", 28}}

  fmt.Println(e)
}
