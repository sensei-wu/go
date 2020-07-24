package main

import (
	"fmt"
	"strings"
)

type strutils struct {
	s string
}

func (str strutils) toUpper() string {
	return strings.ToUpper(str.s)

}

func (str strutils) censor(repl string) string {
	return strings.Repeat(repl, len(str.s))

}

func main() {
	str := new(strutils)
	str.s = "pooru"

	fmt.Println(str.toUpper())

	fmt.Println(str.censor("*"))
}
