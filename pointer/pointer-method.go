package main

import "fmt"

type Man struct {
  Name string
}

func (man *Man) Married() {
  man.Name = "Mr. " + man.Name
  fmt.Println("Di method", man.Name)
}

func main() {
  elang := Man{"Elang"}
  elang.Married()

  fmt.Println(elang.Name)
}

