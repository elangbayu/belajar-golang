package main

import "fmt"

func getFullname() (string, string) {
  return "Elang", "Segara"
}

func main() {
  firstName, lastName := getFullname()
  fmt.Println(firstName, lastName)

  pertama, _ := getFullname()
  fmt.Println(pertama)
}

