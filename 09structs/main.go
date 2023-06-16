package main

import (
	"fmt"
	"reflect"
)

type rectangle struct {
	length  int
	breadth int
	color   string
}

func main() {

	pointers()

	var rect1 = rectangle{10, 20, "Green"}
	fmt.Println(rect1)

	var rect2 = rectangle{length: 10, color: "Green"} // breadth value skipped
	fmt.Println(rect2)

	rect3 := rectangle{10, 20, "Green"}
	fmt.Println(rect3)

	rect4 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(rect4)

	rect5 := rectangle{breadth: 20, color: "Green"} // length value skipped
	fmt.Println(rect5)

	// rect6 is a pointer to an instance of rectangle
	rect6 := new(rectangle) 
	rect6.length = 10
	rect6.breadth = 20
	rect6.color = "Green"
	fmt.Println(rect6)

	// rect7 is an instance of rectangle
	var rect7 = new(rectangle) 
	rect7.length = 10
	rect7.color = "Red"
	fmt.Println(rect7)

	// Struct Instantiation Using Pointer Address Operator
	var rect8 = &rectangle{}
	(*rect8).breadth = 10
	(*rect8).color = "Blue"
	fmt.Println(rect8) // length skipped

	// Find Type of Struct
	fmt.Println(reflect.TypeOf(rect5).Kind())
	fmt.Println(reflect.ValueOf(rect8).Kind())

}

var pointers = func() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}
