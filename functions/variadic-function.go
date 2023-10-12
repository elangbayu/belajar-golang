package main

import "fmt"

func sum(numbers ...int) int {
  total := 0
  for _, value := range numbers {
    total += value
  } 

  return total
}

func main() {
  total := sum(1,2,3,4,5)
  fmt.Println(total)

  slice := []int{1,2,3,4,5,6}
  total = sum(slice...)
  fmt.Println(total)
}

