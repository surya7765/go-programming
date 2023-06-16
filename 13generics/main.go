package main

import (
	"fmt"
)

// Either any or interface{} can be passed
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}



func main() {
	Print([]string{"Hello, ", "playground\n"})
	Print([]int{1, 2, 3})

	var num,num1 float64 = 4.5,5.6
	fmt.Print(multiply(num,num1))

	// type inference
	type float float64
	var num2,num3 float = 4.5,5.6
	fmt.Print(multiply[float](num2,num3))

}

// Type set
type Number interface {
  int|float64|float32
}
func multiply[T ~float64](a,b T) T{
	return a * b
}
