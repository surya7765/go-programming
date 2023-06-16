package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("Hello from GO!")
	content := "Hello from GO!"

	// Creation of the file
	file, err := os.Create("./file.txt")

	if err != nil {
		panic(err)
	}

	// Writing into the file
	Len, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Print("Length: ", Len)

	// Closing the file
	// Better way to close the file is by using defer keyword
	file.Close()
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	fmt.Print(data)
}
