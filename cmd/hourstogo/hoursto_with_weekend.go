package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UTC()
	end := time.Date(2020, time.May, 4, 9, 0, 0, 0, time.UTC)
	//end := start.AddDate(0,0,2)
	diff := end.Sub(start)
	fmt.Println(diff.Hours())
}
