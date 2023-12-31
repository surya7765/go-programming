package main

import (
	"fmt"
	"time"
)

func main() {
	// IF else

	var salary int = 100

	if salary < 50 {
		fmt.Println("you are underpaid")
	} else if salary >= 50 {
		fmt.Println("you are sufficiently paid")
	} else {
		fmt.Println("you are overpaid")
	}

	// https://www.codequizzes.com/golang/beginner/printing-variables

	for ctr := 0; ctr < 10; ctr++ {
		fmt.Println(ctr)
	}

	var num int

	fmt.Scan(&num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%v * %v = %v\n", num, i, num*i)
	}

	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
		}
	}

	city := "Kolkata"
	switch city {
	case "Kolkata":
		fmt.Println("Welcome    to  Kolkata")
	case "Bangalore":
		fmt.Println("Welcome    to  Bangalore")
	case "Mumbai":
		fmt.Println("Welcome Mumbai")
	}

	fmt.Scan(&num)

	switch num {
	case 1, 2, 3:
		fmt.Println("Winter")
	case 4, 5, 6:
		fmt.Println("Sunny")
	case 7, 8, 9:
		fmt.Println("Rainy")
	}

	today := time.Now()
	var t int = today.Day()

	switch t {
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
