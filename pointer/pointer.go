package main

import "fmt"

type Address struct {
  City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
  address.Country = "Indonesia"
}

func main() {
  address1 := Address{"Yogya", "Yogya", "Indonesia"}
  address2 := &address1

  address2.City = "Mojokerto"

  *address2 = Address{"Yogya", "Malang", "Indonesia"}

  fmt.Println(address1)
  fmt.Println(address2)

  var address3 *Address = new(Address)
  fmt.Println(address3)

  var alamat = Address{
    City: "Yogya",
    Province: "Yogya",
    Country: "",
  }
  changeCountryToIndonesia(&alamat)
  fmt.Println(alamat)
}

