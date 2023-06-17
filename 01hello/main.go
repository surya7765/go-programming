package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

// Command for displaying the compilation stages
// go build -gcflags=-S main.go

// Command for storing the compilation stages in txt file
// go build -x -work main.go 1> transcript.txt 2>&1

// Initializing the module
// In general
// go mod init github.com/users/<anyname>

// As usual
// go mod init <anyname>

// To run go file
// go run <filename.go>

// To build go file
// go build <filename.go>

// To format the go file
// go fmt <filename.go>

// To work with multiple files in a single workspace
// go work
// go work use <filename1,filename2>

// Declaration of multiple variables
func main() {
	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)
}

// Working with Print, Println, Printf, and Scan functions
// Taking input using Scan function in "i" variable
func main() {
  var i int

  fmt.Print("Type a number: ")
  fmt.Scan(&i)
  fmt.Printf("%v",i)
  fmt.Println("Your number is:", i)
}


// Working with Print, Println, and Scan functions
// Trying to take two arguments and return different values
func main() {
  var i,j int

  fmt.Print("Type two numbers: ")
  fmt.Scan(&i, &j)
  fmt.Println("Your numbers are:", i, "and", j)
}

// Working with Print, Println, and Scanln functions
// Trying to take two arguments and return different values
func main() {
  var i,j int

  fmt.Print("Type two numbers: ")
  fmt.Scanln(&i, &j)
  fmt.Println("Your numbers are:", i, "and", j)
}

func main() {
  var i,j int

  fmt.Print("Type two numbers: ")
  fmt.Scanf("%v %v",&i, &j)
  fmt.Println("Your numbers are:", i, "and", j)
}

func main(){
    fmt.Println("Hello from Go!")

    var err error
    
    reader := bufio.NewReader(os.Stdin);

    input, err := reader.ReadString(56);
    if err != nil {
        fmt.Println("Error - ", err)
    } else{
        fmt.Println("You entered - ", input)
    }                                                                                                                    
}


func main() {
    
    var num = 30
    
    fmt.Printf("%v\n",num)
    fmt.Printf("%f\n",num)
    fmt.Printf("%T\n",num)
    fmt.Printf("%b\n",num)
    fmt.Printf("%x\n",num)
    fmt.Printf("%X\n",num)
    fmt.Printf("%e\n",num)
    fmt.Printf("%g\n",num)
}



func main() {
    var num float64 = 0.14
    var num1 = 20
    num2 := 40
    fmt.Print(num, num1, num2)
}

func main()  {	
	var num int
	var num1,num2 int
	fmt.Scan(&num, &num1)
	fmt.Scanf("enter any number %v",&num2)
	fmt.Print(num + num1)
	fmt.Println(num * num1)
	fmt.Println(num - num1)
	fmt.Printf("The number is %v",num2)
	fmt.Scanln("%v %v", &num, &num1)
	
}



// Short Variable Declaration(Volorus)
func main() {
	name := "John Doe"
	fmt.Println(reflect.TypeOf(name))
}

// Zero Values variables 
func main() {
	var quantity float32
	fmt.Println(quantity)

	var price int16
	fmt.Println(price)

	var product string
	fmt.Println(product)

	var inStock bool
	fmt.Println(inStock)
}































