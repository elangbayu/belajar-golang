package main

import "fmt"

func main() {
  person := map[string]string {
    "name": "Elang",
    "address": "Yogyakarta",
  }

  person["title"] = "Software Engineer"

  fmt.Println(person)
  fmt.Println(person["name"])
  fmt.Println(person["address"])

  book := make(map[string]string)
  book["title"] = "Belajar golang"
  book["author"] = "Elang"
  book["ups"] = "Salah"

  fmt.Println(book)
  fmt.Println(len(book))

  delete(book, "ups")

  fmt.Println(book)
  fmt.Println(len(book))
}

