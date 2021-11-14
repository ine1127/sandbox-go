package main

import (
  "fmt"
  "strconv"
  "strings"
)

// Data is interface
type Data interface {
  Init(config map[string]string)
  PutData(data map[string]string)
  PrintData()
}

// Student is structure
type Student struct {
  Name string
  Age int
}

func(s *Student) Init(config map[string]string) {
  s.Name = config["name"]
  s.Age, _ = strconv.Atoi(config["age"])
}

func (s *Student) PutData(config map[string]string) {
  s.Name = config["name"]
  s.Age, _ = strconv.Atoi(config["age"])
}

func (s *Student) PrintData() {
  fmt.Printf("Name: %s, Age: %d\n", s.Name, s.Age)
}

// Teacher is structure
type Teacher struct {
  Name string
  Age int
  Subject []string
}

func (s *Teacher) Init(config map[string]string) {
  s.Name = config["name"]
  s.Age, _ = strconv.Atoi(config["age"])
  subject := strings.Split(config["subject"], ",")

  dar := []string{}
  for _, d := range subject {
    dar = append(dar, d)
  }
  s.Subject = dar
}

func (s *Teacher) PutData(config map[string]string) {
  s.Name = config["name"]
  s.Age, _ = strconv.Atoi(config["age"])
  subject := strings.Split(config["subject"], ",")

  dar := []string{}
  for _, d := range subject {
    dar = append(dar, d)
  }
  s.Subject = dar
}

func (s *Teacher) PrintData() {
  fmt.Printf("Name: %s, Age: %d, Subject: %s\n", s.Name, s.Age, s.Subject)
}

func main() {
  var taro Data = new(Student)
  taro.Init(map[string]string{"name": "Tanaka Taro", "age": "11"})
  taro.PrintData()

  var sato Data = new(Teacher)
  sato.Init(map[string]string{"name": "Sato Shinji", "age": "34", "subject": "History,World History"})
  sato.PrintData()
}
