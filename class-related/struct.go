package main

import "fmt"

type Customer struct {
  Name, Address string
  Age int
}

func (customer Customer) sayHello(name string) {
  fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
  var elang Customer
  elang.Name = "Elang"
  elang.Address = "Yogya"
  elang.Age = 24

  elang.sayHello("Segara")

  fmt.Println(elang.Name)
  fmt.Println(elang.Address)
  fmt.Println(elang.Age)

  bayu := Customer{
    Name: "Elang",
    Address: "Yogya",
    Age: 24,
  }
  fmt.Println(bayu)

  segara := Customer{"Elang", "Yogya", 24}
  fmt.Println(segara)
}

