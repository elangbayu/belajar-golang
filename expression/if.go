package main

import "fmt"

func main() {
  name := "Wiwit"

  if name == "Elang" {
    fmt.Println("Hello Elang")
  } else if name == "Wiwit" {
    fmt.Println("Hello Wiwit")
  }else {
    fmt.Println("Hello Guest")
  }

  if length := len(name); length > 10 {
    fmt.Println("Terlalu panjang")
  } else {
    fmt.Println("Valid")
  }
}

