package main

import "fmt"

func main() {

	// Create Empty Slice
	var cities []string

	// Declare Slice using Make
	// Syntax of make function
	// cities := make([]type, len,capacity)\
	// var intSlice = make([]int, 10)        // when length and capacity is same
	// var strSlice = make([]string, 10, 20) // when length and capacity is different

	cities = make([]string, 3)
	
	// Initialize Slice with values using a Slice Literal
	cities = []string{"India", "Canada", "Japan"}

	// Editing cities
	cities[0] = "Kolkata"

	// Appending values to cities
	cities = append(cities, "Amsterdam")
	cities = append(cities, "Den Haag")

	// Accessing items of slices
	fmt.Println(cities[0:2])
	fmt.Println(cities[:3])
	fmt.Println(cities[2:])

	// Remove Item from Slice
	citySlice := RemoveIndex(cities, 3)
	fmt.Println(citySlice)

}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}