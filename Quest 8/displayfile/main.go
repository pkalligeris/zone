package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if exactly one argument is provided
	if len(os.Args) != 2 {
		if len(os.Args) == 1 {
			fmt.Println("File name missing")
			return
		}
		fmt.Println("Too many arguments")
		return
	}

	// Get the filename from arguments
	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename) // tries to open the file in read-only mode.
	if err != nil {                // checks if opening the file failed (e.g., file doesnt exist).
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // ensures the file will be closed automatically when main finishes.

	// Read and print the file contents
	_, err = io.Copy(os.Stdout, file) //_, err = io.Copy(os.Stdout, file)
	if err != nil {                   // If an error occurs while reading, it prints an error message.
		fmt.Println("Error reading file:", err)
		return
	}

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}
