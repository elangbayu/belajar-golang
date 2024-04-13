package main

import "fmt"

func main() {
	// Model pertama
	var person map[string]string = map[string]string{}
	person["name"] = "John Doe"
	person["address"] = "Yogyakarta"

	// Model kedua
	person2 := map[string]string{
		"name":    "John Doe",
		"address": "Yogyakarta",
	}

	fmt.Println(person2["name"])
	fmt.Println(person2["address"])
	fmt.Println(person2)

	fmt.Println(person["salah"]) // jika key tidak ditemukan, maka akan return default value sesuai tipe data

	book := make(map[string]string)
	book["title"] = "The Go Programming Language"
	book["author"] = "Go Author"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
