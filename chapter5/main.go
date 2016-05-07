package main

import "fmt"

func main() {
  var x [5]int
  x[4] = 100
  fmt.Println(x)

  var y [5]float64
  y[0] = 98
  y[1] = 93
  y[2] = 77
  y[3] = 82
  y[4] = 83

  total := 0.0
  // var total float64 = 0
  for i:= 0; i < 5; i++ {
    total += y[i]
  }
  fmt.Println(total / float64(len(y)))

  total = 0
  for _, value := range y {
    total += value
  }
  fmt.Println(total / float64(len(y)))

  z := make([]float64, 5)
  fmt.Println(z)

  slice1 := []int{1,2,3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)

  mmap := make(map[string]int)
  mmap["key"] = 10
  mmap["anotherKey"] = 42
  fmt.Println(mmap)


  // elements := make(map[string]string)
  // elements["H"] = "Hydrogen"
  // elements["He"] = "Helium"
  // elements["Li"] = "Lithium"
  // elements["Be"] = "Beryllium"
  // elements["B"] = "Boron"
  // elements["C"] = "Carbon"
  // elements["N"] = "Nitrogen"
  // elements["O"] = "Oxygen"
  // elements["F"] = "Fluorine"
  // elements["Ne"] = "Neon"

  elements := map[string]string {
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
    "Be": "Beryllium",
    "B": "Boron",
    "C": "Carbon",
    "N": "Nitrogen",
    "O": "Oxygen",
    "F": "Fluorine",
    "Ne": "Neon",
  }

  fmt.Println(elements["Li"])
  fmt.Println(elements["Un"])

  if name, ok := elements["Un"]; ok  {
    fmt.Println(name, ok)
  }

  elems := map[string]map[string]string {
    "H": map[string]string {
      "name": "Hydrogen",
      "state": "gas",
    },
    "He": map[string]string {
      "name": "Helium",
      "state": "gas",
    },
    "Li": map[string]string {
      "name": "Lithium",
      "state": "solid",
    },
    "Be": map[string]string {
      "name": "Beryllium",
      "state": "solid",
    },
    "B": map[string]string {
      "name": "Boron",
      "state": "solid",
    },
    "C": map[string]string {
      "name": "Carbon",
      "state": "solid",
    },
    "N": map[string]string {
      "name": "Nitrogen",
      "state": "gas",
    },
    "O": map[string]string {
      "name": "Oxygen",
      "state": "gas",
    },
    "F": map[string]string {
      "name": "Fluorine",
      "state": "gas",
    },
    "Ne": map[string]string {
      "name": "Neon",
      "state": "gas",
    },
  }

  if el, ok := elems["Li"]; ok {
    fmt.Println(el["name"], el["state"])
  }

  tmp := make([]int, 3, 9)
  fmt.Println(len(tmp))

  tmp2 := [6]int{0,1,2,3,4,5}
  fmt.Println(tmp2[2:5])

  arr := []int{
    48, 96, 86, 68,
    57, 82, 63, 70,
    37, 34, 83, 27,
    19, 97,  9, 17,
  }

  smallestValue := arr[0]
  for _, currentValue := range arr {
    if smallestValue > currentValue {
      smallestValue = currentValue
    }
  }
  fmt.Println(smallestValue)
}