package main

import (
	"fmt"
)

func main() {

	doSomething()

	sum := addValues(5, 8)

	fmt.Println("The sum is:", sum)

	multiSum, multiCount := addAllValues(4, 5, 987)

	fmt.Println("Sum of multiple values:", multiSum)
	//If you delcare a function that returns multiple values, make sure to declare another variable to store the second value in it:
	fmt.Println("Sum of multiple values:", multiCount)

}

// if you do not need to return anything in your function, you do not need anything between the parens and the braces:
func doSomething() {
	fmt.Println("Do something")
}

//If the function is going to return a value, declare the value after the functions closing parentheses.  If you are wanting to pass in any arguments
//, do so by declaring them inside the parentheses.  Also, if the func is accepting multiple arguments of the same type, you only need to declare the type for the last argument in the parentheses:

func addValues(value1, value2 int) int {
	return value1 + value2
}

//a func can also accept an arbitrary number of values of the same type.  Go also lets you return multiple values from a function
//Go to the function, change the signiture so that you are returning a comma, delimited list wrapped in parentheses:

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}

	return total, len(values)
}
