package main

import (
	hash "double-gopher/internal/Hash"
	storage "double-gopher/internal/Storage"
	"fmt"
	"os"
)

func main() {
	//Variable to keep user input
	var input string

	//Print the invite for the input
	fmt.Println("Input the folder in which you want to get out of the doubles !!! :) ")

	//Assign user input to variable
	fmt.Scanln(&input)

	//Get the fs.FS of the choosen directory
	fsys := os.DirFS(input)

	//Change working directory
	os.Chdir(input)

	// Create new storage
	storage := storage.NewStorage()

	//Run the code
	hash.HashFiles(fsys, storage)

}
