package main

import "fmt"


func main() {
	fmt.Println("Hello from Go!")

	var input string
	var err error

	_, err = fmt.Scanln(&input)

	if err != nil {
		defer First()
		panic(err)
	} else {
		fmt.Println("You entered - ", input)
	}
}

func First() {
	panic("unimplemented")
}
