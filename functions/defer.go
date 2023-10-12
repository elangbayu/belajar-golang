package main

import "fmt"

func logging() {
  fmt.Println("Welcome to dashboard")
}

func runApp(value int) {
  defer logging()
  fmt.Println("Hello Guest")
  result := 10 / value
  fmt.Println("Result", result)
}

func main() {
  runApp(10)
}

