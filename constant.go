package main

import "fmt"

func main() {
	// Model pertama
	const firstName string = "Elang"
	const lastName = "Segara"
	// constant tidak masalah jika tidak dipakai

	// error: constant tidak bisa diubah lagi setelah dideklarasikan
	// firstName = "Bayu" //error
	// lastName = "Wiwit" //error

	// Model kedua
	const (
		firstName1 string = "Elang"
		lastName1         = "Segara"
	)

	fmt.Println(firstName1)
	fmt.Println(lastName1)
}
