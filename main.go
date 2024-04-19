package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "Hello"

	//we are now creating a variable to store the newly created file in.  We are also handling an error, so we create a variable for that too.
	//use the os package and the Create(), and we are using ./ to reference the same file path as the current .go file we are working with
	//we are also passing our err variable into the checkError() we made earlier
	file, err := os.Create("./fromString.txt")
	checkError(err)
	//Now, we are creating another variable called length, and another error handling variable, err.  we will be using the io package with the
	//WriteString() that takes two arguments, a writer object and a string. We are passing in the file variable, and the content variable.
	//We check the err again with the checkError func
	//the length variable will be returned, and it will be the number of characters added to the file.
	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Wrote a file with %v characters\n", length)
	//defer means to wait until everything else is done, and then execute the command.  We always want to make sure we close a file after altering
	//it
	defer file.Close()

	defer readFile("./fromString.txt")

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//When you read a file, it always comes to you as an array of bytes.

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))

}
