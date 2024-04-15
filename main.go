package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Initializing the variable with bufio.NewReader(), this function requires a parameter indicating where the data is going to come from, in this case, it will come from os.Stdin.  This stands for Standard Input, what the user types on the command line.
	reader := bufio.NewReader(os.Stdin)

	//Displaying a prompt to the reader:
	fmt.Print("Enter text: ")

	//Now, we want the application to pause and let the user type something in.
	//We are declaring two variables, input for the input that the user will be typing into the console, and _, for error handling.  _ means that we want to ignore that variable.
	//We are initializing the variables with reader.ReadString(), this function accepts a single character, the deliminator, which tells the ReadString function when to stop receiving input.  In this case, we are inputing the argument of '\n' to print a new line as a character
	//The ReadString function knows to stop reading the string that was entered after the new string that gets printed by the Println().  Print() doesn't print a new line after the argument is passed in.  But, Println does.
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")

	//we are initializing another variable using the bufio package.  We are reusing reader from last time, but now for numInputm, and we are ignoring the error again.
	numInput, _ := reader.ReadString('\n')

	//Now, we are converting the string to an integer, and we want to pay attention to the error.  The user could input a value that cannot be converted to an integer value.
	//we are converting the string to a Float, and we are using the strconv package, and the ParseFloat function, which takes in two parameters: the string that we are converting, and the bit size.
	//We want to handle for if the user adds any spaces before or after the input, so we use the TrimSpace function and insert the numInput as the parameter to remove any trailing whitespace.
	//Then, we input 64 because we are using a 64 bit operating system.
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	//Now, we are handling for the error.  If the error variable is not nil (null), meaning the error takes on the value of the invalid input by the user; then, we print the error message to the screen.
	//If the error message is nil, we will print the value to the screen by calling the converted Float value
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", aFloat)
	}
}
