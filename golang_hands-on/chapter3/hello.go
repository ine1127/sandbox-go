package main

import (
  "fmt"
)

// Data is interface
type Data interface {
  Initial(name string, data []int)
  PrintData()
}

// Mydata is Struct.
type Mydata struct {
  Name string
  Data []int
}

// Initial is init method
func (md *Mydata) Initial(name string, data []int) {
  md.Name = name
  md.Data = data
}

// PrintData is println all data.
func (md *Mydata) PrintData() {
  fmt.Println("Name: ", md.Name)
  fmt.Println("Data: ", md.Data)
}

func main() {
  var ob Data = new(Mydata)
  ob.Initial("Sachiko", []int{55, 66, 77})
  ob.PrintData()
}
