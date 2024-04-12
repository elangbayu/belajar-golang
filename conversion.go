package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // number overflow, karena 32768 lebih dari kapasitas int16, maka akan mengulang dari range terbawah yaitu -32768

	var name = "Elang Segara"
	var e uint8 = name[0]
	var eString = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
