package main

import (
	"fmt"
	"time"
)

func main() {
	jam := time.Now()

	fmt.Println(jam)
	fmt.Println(jam.Year())
	fmt.Println(jam.Month())
	fmt.Println(jam.Date())
	fmt.Println(jam.Hour())
	fmt.Println(jam.Minute())
	fmt.Println(jam.Second())
	fmt.Println(jam.Nanosecond())

	//klo mau membuat tanggal sendiri
	clock := time.Date(2023, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(clock)

	//parsing tanggal
	layout := "2006-09-20"
	parsingClock, _ := time.Parse(layout, "2020-10-20")
	fmt.Println(parsingClock)
}
