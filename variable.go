package main

import "fmt"

func main() {
	// Model pertama
	var name string

	name = "Elang Bayu Segara"
	fmt.Println(name)

	name = "Elang Segara"
	fmt.Println(name)

	// Model kedua
	var name2 = "Elang Bayu Segara" // bisa juga seperti ini: var name2 string = "Elang Bayu Segara"
	fmt.Println(name2)

	name2 = "Elang Segara"
	fmt.Println(name2)

	// Model ketiga
	name3 := "Elang Bayu Segara"
	fmt.Println(name3)

	name3 = "Elang Segara"
	fmt.Println(name3)

	// Model ketiga
	var (
		firstName  = "Elang"
		middleName = "Bayu"
		lastName   = "Segara"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
