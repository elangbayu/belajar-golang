package main

import (
  "belajar-golang/db"
  "fmt"
)

func main() {
  result := db.GetDatabase()
  fmt.Println(result)
}

