package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var jaba string
	var x int
	var b bool

	fmt.Printf("Size of %s is %d\n", jaba, unsafe.Sizeof(jaba))
	fmt.Printf("Size of %d is %d\n", x, unsafe.Sizeof(x))
	fmt.Printf("Size of %v is %d\n", b, unsafe.Sizeof(b))
}