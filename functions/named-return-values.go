package main

import "fmt"

func getFullname2() (firstName string, middleName string, lastName string) {
  firstName = "Elang"
  middleName = "Bayu"
  lastName = "Segara"

  return
}

func main() {
  firstName, middleName, lastName := getFullname2()
  fmt.Println(firstName)
  fmt.Println(middleName)
  fmt.Println(lastName)
}
