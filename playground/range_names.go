package main

import "fmt"

func main() {

  names := [3]string{"Aidan", "Fanny", "Delina"}

  for _, name := range names {
    fmt.Println(name)
  }
}
