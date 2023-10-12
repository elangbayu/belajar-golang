package main

import "fmt"

type HasName interface {
  GetName() string
}

func sayHello(name HasName) {
  fmt.Println("Hello", name.GetName())
}

type Person struct {
  Name string
}

func (person Person) GetName() string {
  return person.Name
}

func main() {
  var elang Person
  elang.Name = "Elang"

  sayHello(elang)
}

