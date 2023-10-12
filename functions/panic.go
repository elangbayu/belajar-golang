package main

import "fmt"

func endApp() {
  fmt.Println("Done")
  message := recover()
  fmt.Println("Error:", message)
}

func runApp(error bool) {
  defer endApp()
  if error {
    panic("ERROR")
  }
  fmt.Println("Running")
}

func main() {
  runApp(false)
}

