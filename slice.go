package main

import "fmt"

func main() {
	names := [...]string{"John", "Paul", "George", "Ringo", "Elang", "Segara"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:] // jika dideklarasikan manual akan seperti var slice4 []string = names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "Sabtu Baru" // jika kita ubah slice dari array maka elemen pada array akan ikut berubah karena slice merupakan pointer ke array tersebut
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days) // Sabtu dan Minggu pada array days akan berubah menjadi Sabtu Baru dan Minggu Baru

	daySlice2 := append(daySlice1, "Libur Baru") // append akan membuat slice baru, maka perlu disimpan ke sebuah variable baru
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice2)
	// daySlice1 dan days tidak akan berubah karena daySlice2 sudah bukan pointer ke array days lagi, karena sudah dibuatkan array baru
	fmt.Println(daySlice1)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5) // panjang pertama 2, tapi maksimal kapasitas 5
	newSlice[0] = "Elang"
	newSlice[1] = "Elang"
	//newSlice[2] = "Elang" error karena jika mau menambahkan elemen melebihi panjang array harus menggunakan append()

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Elang")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2)) // capacity masih 5 karena masih menggunakan array yang sama

	newSlice2[0] = "Budi" // jika kita ubah index pertama dari newSlice2, maka newSlice juga akan berubah karena append di atas tidak melebih kapasitas jadi newSlice2 masih menggunakan array yang sama (newSlice2 adalah pointer ke newSlice)
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// CAUTION
	// deklarasi array dan slice berbeda pada square bracketnya, kalau array harus ada isinya meskipun ...
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
