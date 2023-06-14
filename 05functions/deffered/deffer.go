package main

import "fmt"

func First() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}

func main() {
	defer println("World")
	println("Hello")
	defer second()
	first()
}
