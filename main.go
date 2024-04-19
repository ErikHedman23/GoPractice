package main

import (
	"fmt"
	"strconv"
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

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.

	newVal1 := convertInputToValue(input1)
	newVal2 := convertInputToValue(input2)
	var answer float64
	switch operation {
	case "+":
		answer = addValuess(newVal1, newVal2)
	case "-":
		answer = subtractValues(newVal1, newVal2)
	case "*":
		answer = multiplyValues(newVal1, newVal2)
	case "/":
		answer = divideValues(newVal1, newVal2)
	default:
		panic("Invalid operator!")
	}
	return answer
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", input)
		panic(message)
	} else {
		return value
	}
}

func addValuess(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
