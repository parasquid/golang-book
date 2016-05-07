package main

import "fmt"

func main() {
  for i:=1; i <= 10; i++ {
    if i % 2 == 0 {
      fmt.Println(i, "even")
    } else {
      fmt.Println(i, "odd")
    }
  }

  for j:=0; j <= 10; j++ {
    switch j {
      case 0: fmt.Println("Zero")
      case 1: fmt.Println("One")
      case 2: fmt.Println("Two")
      case 3: fmt.Println("Three")
      case 4: fmt.Println("Four")
      case 5: fmt.Println("Five")
      default: fmt.Println("Unknown Number")
    }
  }

  for k := 1; k <= 100; k++ {
    if k % 3 == 0 && k % 5 == 0 {
      fmt.Println("FizzBuzz")
    } else if k % 3 == 0 {
      fmt.Println("Fizz")
    } else if k % 5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(k)
    }
  }
}