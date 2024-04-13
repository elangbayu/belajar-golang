package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Elang"
	names[1] = "Bayu"
	names[2] = "Segara"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Model deklarasi array kedua
	var values = [3]int{
		90,
		80, // index ke 3 akan terisi oleh default value yaitu 0
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Bisa juga ukuran array mengikuti ukuran valuenya, di sini berarti values2 akan mempunyai length 4
	var values2 = [...]int{
		90,
		80,
		10,
		40,
	}

	fmt.Println(values2)
	fmt.Println(len(values2))
}
