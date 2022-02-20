package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with Text File in Go")

	// our default content
	content := "This is file content that will be saved on a file"

	// create file

	file, err := os.Create("./programTextFile.txt")
	checkNilError(err)

	// write the content to file
	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length of file content: ", length)
	readFile("./programTextFile.txt")

	// close the file after all operation done

	defer file.Close()
}

// write file content

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkNilError(err)

	fmt.Println("File content:\n", string(data))
}

// error handle function to reuse

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
