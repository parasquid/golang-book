package main

import "fmt"

func main() {
  fmt.Println("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  var output float64

  output = (input - 32) * 5/9
  fmt.Println(output)

  output = input * 0.3048
  fmt.Println(output)
}