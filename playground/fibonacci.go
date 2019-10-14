package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  n, _ := strconv.Atoi(os.Args[1])
  fibonacci(n)
}

func fibonacci(n int) {
  last := 1
  lastButOne := 0
  val := 0

  fmt.Println(lastButOne)
  fmt.Println(last)

  for i := 2; i < n ; i++ {
      val = last + lastButOne
      fmt.Println(val)
      lastButOne, last = last, val
    }
}
