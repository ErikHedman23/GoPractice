package main

import (
	"bufio"
	"fmt"
	"os"
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
}
