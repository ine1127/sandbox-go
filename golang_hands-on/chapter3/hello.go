package main

import (
  "fmt"
  "strings"
  "strconv"
)

// Data is interface
type Data interface {
  SetValue(vals map[string]string)
  PrintData()
}

// Mydata is structure.
type Mydata struct {
  Name string
  Data []int
}

// SetValue is Mydata method.
func (md *Mydata) SetValue(vals map[string]string) {
  md.Name = vals["name"]
  valt := strings.Split(vals["data"], " ")
  vali := []int{}

  for _, i := range valt {
    n, _ := strconv.Atoi(i)
    vali = append(vali, n)
  }
  md.Data = vali
}

// PrintData is Mydata method.
func (md *Mydata) PrintData() {
  if md != nil {
    fmt.Println("Name: ", md.Name)
    fmt.Println("Data: ", md.Data)
  } else {
    fmt.Println("**This is Nil value.**")
  }
}

// Yourdata is structure.
type Yourdata struct {
  Name string
  Mail string
  Age int
}

// SetValue is Yourdata method.
func (md *Yourdata) SetValue(vals map[string]string) {
  md.Name = vals["name"]
  md.Mail = vals["mail"]
  n, _ := strconv.Atoi(vals["age"])
  md.Age = n
}

// PrintData is Yourdata method.
func (md *Yourdata) PrintData() {
  fmt.Printf("I'm %s. (%d).\n", md.Name, md.Age)
  fmt.Printf("mail: %s.\n", md.Mail)
}

func main() {
  var ob *Mydata
  ob.PrintData()
  ob = &Mydata{}
  ob.SetValue(map[string]string{
          "name": "Jiro",
          "data": "123 456 789",
  })

  ob.PrintData()
}
