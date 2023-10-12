package main

import "fmt"

func getGoodbye(name string) string {
  return "Bye " + name
}

func main() {
  goodbyeMessage := getGoodbye
  result := goodbyeMessage("Elang")
  fmt.Println(result)
}

