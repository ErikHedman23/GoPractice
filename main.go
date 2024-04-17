package main

import (
	"fmt"
)

//variables that store the memory address of another variable

func main() {
	anInt := 42

	var p = &anInt //the & means that you are pointing at the memory address of the variable, not at the value.

	fmt.Println("Value of p:", *p) //*p is pointing at the memory address of p
	fmt.Println("Address of anInt:", &anInt)

	value1 := 42.15
	pointer1 := &value1 //not explicitly declaring the type of pointer1, but by using the &, you are implicitly delcaring pointer1 as a pointer, pointing to the memory address
	//of value1

	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31

	//this changes the original value of value1

	fmt.Println("Pointer 1:", value1)

	//Pointers are great when you are wanting to access and alter the original value of a variable.

}
