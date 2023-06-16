package main

import "fmt"
/*
type Human interface {
	Display()
}

type Employee struct {
	name        string
	designation string
}

func (emp Employee) Display() {
	fmt.Println("Name - ", emp.name, ", Designation - ", emp.designation)
}

type Contractor struct {
	name        string
	weeklyHours int
}

func (cont Contractor) Display() {
	fmt.Println("Name - ", cont.name, ", Weekly Hours - ", cont.weeklyHours)
}

// Employee is an interface for printing employee details
type Employee1 interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

// Emp user-defined type
type Emp int

// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func main() {
	var emp Human = Employee{name: "Rob Pike", designation: "Engineer"}
	emp.Display()
	var cont Human = Contractor{name: "XYZ", weeklyHours: 35}
	cont.Display()

	var e1 Employee1
	e1 = Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}
*/
// Interfaces with common Method

type Vehicle interface {
	Structure() []string // Common Method
	Speed() string
}

type Human interface {
	Structure() []string // Common Method
	Performance() string
}

type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

func (c Car) Speed() string {
	return "200 Km/Hrs"
}

type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return parts
}

func (m Man) Performance() string {
	return "8 Hrs/Day"
}

func main() {
	var bmw Vehicle
	bmw = Car("World Top Brand")

	var labour Human
	labour = Man("Software Developer")

	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}
}
