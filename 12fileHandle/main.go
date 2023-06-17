package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// fmt.Print("Hello from GO!")
	content := "Hello from GO!"

	// Creation of the file
	file, err := os.Create("./file.txt")
	checkForError(err)

	// Writing into the file
	Len, err := io.WriteString(file, content)

	checkForError(err)

	fmt.Println("Length: ", Len)

	readFile(file.Name())

	// Closing the file
	// Better way to close the file is by using defer keyword
	file.Close()

	createFolder()

	renameFile()

	copyFile()

	deleteFile()
}

func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

// Reading the file
func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkForError(err)
	fmt.Println(string(data))
}

func createFolder() {
	_, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		checkForError(errDir)
	} else {
		fmt.Print(err)
	}
}

func renameFile() {
	oldName := "./file.txt"
	newName := "./Gofile.txt"
	err := os.Rename(oldName, newName)
	checkForError(err)
}

func copyFile() {
	sourceFile, err := os.Open("./Gofile.txt")
	checkForError(err)
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("./test.txt")
	checkForError(err)
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	checkForError(err)
	fmt.Printf("Copied %d bytes.", bytesCopied)
}

func deleteFile(){
	err := os.Remove("./test.txt")
	checkForError(err)
}