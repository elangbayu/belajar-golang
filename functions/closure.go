package main

import "fmt"

func main() {
  name := "Elang"
  counter := 0

  increment := func() {
    name = "Segara"
    fmt.Println("Increment")
    counter++
    fmt.Println(name)
  }

  increment()
  increment()

  fmt.Println(counter)
  fmt.Println(name)
}

