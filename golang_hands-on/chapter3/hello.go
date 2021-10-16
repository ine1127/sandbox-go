package main

import (
  "fmt"
)

// Mydata is structure
type Mydata struct {
  Name string
  Data []int
}

// PrintData is println all data.
func (md Mydata) PrintData() {
  fmt.Println("*** Mydata ***")
  fmt.Println("Name: ", md.Name)
  fmt.Println("Data: ", md.Data)
  fmt.Println("*** end ***")
}

func main () {
  hanako := Mydata{
            "Hanako", []int{98, 76, 54, 32, 10},
  }
  hanako.PrintData()
}
