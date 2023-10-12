package main

import "fmt"

func main() {

  for counter := 1; counter <= 10; counter++ {
    fmt.Println("Perulangan ke", counter)
  }

  slice := []string{"Elang", "Bayu", "Segara"}

  for i := 0; i < len(slice); i++ {
    fmt.Println(slice[i])
  }

  for _, value := range slice {
    // fmt.Println("Index", i, "=", value)
    fmt.Println(value)
  }

  person := make(map[string]string)
  person["firstname"] = "Elang"
  person["lastname"] = "Segara"

  for key, value := range person {
    fmt.Println(key, "=", value)
  }
}
