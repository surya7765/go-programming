package main

import "fmt"

func main() {
	cities := [...]string{"Kolkata", "Chennai", "Blore"}
	fmt.Println(cities)
	fmt.Println(len(cities))

	// only array's index
	for index := range cities {
		fmt.Println(cities[index])
	}

	// With array index
	for index, value := range cities {
		fmt.Println(index, value)
	}

	// Only array values
	for _, value := range cities {
		fmt.Println(value)
	}

}
