package main

import (
	"fmt"
)

func main() {
	//This is creating an instance of the struct Dog.  This is like defining a constructor in other languages, but you are not calling a method,
	//you are defining an object, so we do not use (), we use {}
	//Pass in the values of the object in the order they are declared:
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	//to see a structs values for debugging purposes, use the Printf()
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	//if you want to change the values of the fields in the struct:

	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

}

//If you name the type with an upper case initial letter, it can be used globally, or in Go terms, it is exported
//If you name the type with a lower case initial letter, it is considered private, or non exported

//The same goes for the properties (fields) of the struct

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	// each value of a struct is technically a field, you access the structs individual fields using the . operator
}
