package main

import "fmt"


func main() {
	fmt.Println("Hello from Go")
	var result int = add(20, 10)
	fmt.Println(result)

	var x, y int = 10, 20
	var p, q int = swap(x, y)
	fmt.Print(p, q)

	fmt.Print("Area: ", rectangle(4, 5))
}

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)

	area = l * b
	return // Return statement without specify variable name
}
