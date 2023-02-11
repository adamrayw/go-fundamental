package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow)

	timeCustom := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(timeCustom)

	layout := "2006-01-02 15:04:05"
	fmt.Println(timeNow.Format(layout))
}
