package main

import "fmt"

// SimpleFunction prints a message
func SimpleFunction() {
	fmt.Println("Hello World")
}

func main() {
	SimpleFunction()
	add(2,3)
}



// Function accepting arguments
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)

	area = l * b
	return // Return statement without specify variable name
}


/*
Golang Functions Returning Multiple Values

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func main() {
	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
}
*/

// anonymus function
func(){
  fmt.Println("Welcome")
}()

 // Passing arguments in anonymous function
func(ele string){
	fmt.Println(ele)
}("hello")



// To define different function in a single line
var (
	area = func(l int, b int) int {
		return l * b
	}
	area1 = func(l int, b int) int {
		return l * b
	}
)



