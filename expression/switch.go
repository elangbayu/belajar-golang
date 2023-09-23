package main

import "fmt"

func main() {
  name := "Wiwit"

  switch name {
  case "Elang":
    fmt.Println("Hello Elang")
  case "Segara":
    fmt.Println("Hello Segara")
  default:
    fmt.Println("Hello Guest")
  }

  switch length := len(name); length > 5 {
  case true:
    fmt.Println("Terlalu panjang")
  case false:
    fmt.Println("Nama sudah benar")
  }

  switch length := len(name); {
  case length > 5:
    fmt.Println("Terlalu panjang")
  case length > 3:
    fmt.Println("Cukup panjang")
  default:
    fmt.Println("Valid")
  }
}

