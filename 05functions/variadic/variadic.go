package main

import "fmt"

func main() {
	displayMessage("Call1", 1)
	displayMessage("Call2", 2, "Param1")
	displayMessage("Call3", 3, "Param1", "Param2", "Param3", "Param4")

}

//variadic functions
func displayMessage(message string, times int, params ...string) {
	fmt.Println(message, times, params)
}
