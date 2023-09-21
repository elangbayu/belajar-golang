package main

import "fmt"

func main() {
  type NoKtp string
  type Married bool

  var noKtpElang NoKtp = "123456789"
  var marriageStatus Married = false
  fmt.Println(noKtpElang)
  fmt.Println(marriageStatus)
}

