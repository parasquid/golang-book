package main

import "fmt"

func average(xs []float64) (result float64) {
  total := 0.0
  for _, v := range xs {
    total += v
  }
  result = total / float64(len(xs))
  return
}

func makeEvenGenerator() (func() uint) {
  i := uint(0)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func makeOddGenerator() (func() uint) {
  i := uint(1)
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

func fib(n int) int {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  }
  return fib(n - 1) + fib(n - 2)
}

func factorial(x uint) uint {
  if x == 0 {
    return 1
  }
  return x * factorial(x - 1)
}

func first() {
  fmt.Println("1st")
}

func second() {
  fmt.Println("2nd")
}

func summer(slice []int) int {
  total := 0
  for _, value := range slice {
    total += value
  }
  return total
}

func half(num int) (int, bool) {
  halved := num / 2
  even := false
  if num % 2 == 0 {
    even = true
  }
  return halved, even
}

func findBiggest(nums ...int) int {
  biggest := nums[0]
  for _, value := range nums {
    if value > biggest {
      biggest = value
    }
  }
  return biggest
}

func main() {
  numbers := []int{1,2,3}
  fmt.Println(summer(numbers))

  fmt.Println(half(1))
  fmt.Println(half(2))

  fmt.Println(findBiggest(1,2,3,4,5,4))

  fmt.Println("fib", fib(7))

  xs := []float64{98, 93, 77, 82, 83}
  fmt.Println(average(xs))

  add := func(x, y int) int {
    return x + y
  }
  fmt.Println(add(1,1))

  x := 0
  increment := func() int {
    x++
    return x
  }
  fmt.Println(increment())
  fmt.Println(increment())

  nextEven := makeEvenGenerator()
  fmt.Println(nextEven())
  fmt.Println(nextEven())
  fmt.Println(nextEven())

  defer second()
  first()

  defer func() {
    str := recover()
    fmt.Println(str)
  }()
  panic("PANIC")
}