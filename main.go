package main

import (
	"fmt"
)

// since the variable starts with a lowercase a, it is private and only can be accessed within this file.
// If it had an uppercase A at the beginning, it would be public, and it would be accessible within any file within this directory
const aConst int = 64

func main() {

	var aString string = "This is Go!"

	fmt.Println(aString)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("This variable's type is %T\n", defaultInt)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("This variable's type is %T\n", anotherInt)

	//another way of implicitly typing a variable is by using :=
	//You can only use the := to implicitly type variables within a function.  Any variable declared outsite of a function must use the var keyword.
	//Same for constants, you must use the const keyword
	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("This variable's type is %T\n", myString)

}
