package main

import (
	"fmt"
)

func main() {

	// Empty map initialization
	var employee = map[string]int{}
	// Initialization of map by default
	employee = map[string]int{"Mark": 10, "Sandy": 20}

	// Initialization of map using make function
	emps := make(map[string]string)

	// Adding Items to map
	emps["bill12"] = "Bill"
	emps["satya342"] = "Satya"
	emps["sundar"] = "Sunder"
	emps["tatte"] = "Andrew"

	// Editing map
	employee["Mark"] = 50 // Edit item

	// Printing map after editing
	fmt.Print(employee)

	// Deleting map
	delete(employee, "Mark")

	// Printing map after deleting
	fmt.Print(employee)

	// Printing map with keys and values pairings
	fmt.Println(emps)
	// Printing length of map
	fmt.Println(len(emps))

	// Iterate over a Map
	for key, element := range employee {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	// Truncate Map
	// Method - I
	for k := range employee {
		delete(employee, k)
	}
	// Method - II
	employee = make(map[string]int)


	

}
