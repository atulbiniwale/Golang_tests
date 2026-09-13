package main

import (
	"fmt"
	"io"
	"os"
)

// file Handling (creating and writing content in file) > close a file

func main() {

	// file creation and closing.
	file, err := os.Create("samplefile.txt") // create a file using os module.

	if err != nil {
		fmt.Println("error while creating file", err)
	}
	defer file.Close() // always close the file.

	// adding/writing contents to the file.
	content := "writing sample text in the file"
	_, err2 := io.WriteString(file, content)
	if err2 != nil {
		fmt.Println("errror while writing content to file", err2)
		return
	}
	fmt.Println("successfully written in file.")

	// reading file contents into the buffer.
	data, _ := os.ReadFile("samplefile.txt")
	fmt.Println(string(data))

}
