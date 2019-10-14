package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UTC()
	end := time.Date(2019, time.December, 18, 0, 0, 0, 0, time.UTC)
  //end := start.AddDate(0,0,2)

  days := 0

  for {

    if(start.After(end)) {
      break
    }
    if(start.Weekday() != 6 && start.Weekday() != 0) {
      days++
    }

    start = start.Add(time.Hour*24)

  }

	fmt.Println(days)
}
