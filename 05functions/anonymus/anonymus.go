package main

import (
	"fmt"
)

var sub = func(x int, y int) int {
	return x - y
}

func main() {
	//Add is declared as anonymus function
	// add := func(x int, y int) int {
	// 	return x + y
	// }
	// fmt.Println(add(20, 10)) //30
	fmt.Println(sub(20, 10)) //10

	for i := 10.0; i < 100; i += 10.0 {
		// Closure function
		rad := func() float64 { //defination of function
			return i * 39.370
		}() //calling function

		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}

	value := func(p, q string) string {
		return p + " " + q
	}
	printHello(value) // printing the function

	add := func(x int, y int) int {
		return x + y
	}

	sub := func(x int, y int) int {
		return x - y
	}
	result := doWork(add, 30, 20)
	fmt.Println(result)
	result = doWork(sub, 30, 20)
	fmt.Println(result)

}
// Higher Order Functions 
// Passing anonymous function
// as an argument
func printHello(fnName func(p, q string) string) {
	fmt.Println(fnName("Hello", "World"))
}

type HigherFunc func(x int, y int) int	// User defined function type
func doWork(anonymous HigherFunc, num1 int, num2 int) int {
	return anonymous(num1*num1, num2*num2)
}
