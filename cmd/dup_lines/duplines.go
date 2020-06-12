package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	count := make(map[string] int)

	for scanner.Scan() {

		line := scanner.Text()

		count[line]++

		if count[line] > 1 {
			fmt.Printf("%s %s\t%d\n", "Found duplicate:", line, count[line])
		}
	}

}
