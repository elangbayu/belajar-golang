package main

import "fmt"

type Blacklist func(string) bool

func regist(name string, blacklist Blacklist) {
  if blacklist(name) {
    fmt.Println(name, "is blocked")
  } else {
    fmt.Println("Welcome", name)
  }
}

func main() {
  blacklist := func(name string) bool {
    return name == "admin"
  }

  regist("admin", blacklist)
  regist("Elang", blacklist)
}

