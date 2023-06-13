package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()

	fmt.Println(presentTime)

	// To format time in GO this is the restricted time format syntax

	fmt.Println(presentTime.Format("15:04:05"))

	fmt.Println(presentTime.Format("Monday"))

	fmt.Println(presentTime.Format("2006-01-02"))

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	time.Sleep(time.Second * 2)
}
