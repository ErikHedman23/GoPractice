package main

import (
	"fmt"
)

type cartItem struct {
	name     string
	price    float64
	quantity int
}

func calculateTotal(cart []cartItem) float64 {
	var value float64 = 0
	for _, item := range cart {
		value += (item.price * float64(item.quantity))
	}
	return value
}

func main() {

	colors := []string{"Red", "Green", "Blue"}
	//Go has several iterations of the for loop.  It does not have a while loop, but the for loop can act as a while loop when needed.
	//To perform the traditional C style for loop with 3 declarations:
	//Here, you are initializing a local variable of i for the index of the slice starting at 0
	//Then, you are stating that you will continue the loop till you get to the end of the length of the slice of i
	//Finally, you are incrementing by 1 index in the slice each time using the incrementor of i++
	//For each time the loop iterates over the slice, we print to the console the index of colors using colors[i]
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	//This style of for loop is the exact same as the one above, it is just easier to read and more manageable:
	for i := range colors {
		fmt.Println(colors[i])
	}

	/*The next way of using a for loop is like how you would use a foreach loop, although slightly different;
		You will be returning two variables from the range from a comma delimited list.
		The first variable will represent the index, and if you want to ignore the index, you can simply put the _ to ignore it.
		The second variable will be the value of the index
	It works just like how it does in C#, only you also get to return the index if you wish:
	for(int color in colors)
	{

	}
	*/
	for _, color := range colors {
		fmt.Println(color)
	}

	//In Go, there is no while loop, but you can initialize a counter or a boolean expression to act as your condition for how long the loop
	//iterates for:
	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	//break and continue are used in Go for for loops.  break means to go to the end of the current code block, this works for both for and
	//switch statements.  continue means to start back at the beginning of the current code block
	//There is also the goto statement with labels.  You can add labels to your code and jump to your label in your code with a goto or a break
	//statement.
	//You can do the same thing with this by using a break, you just wouldn't create a label, you would just say break, and then once the code hops out
	//of the for loop, you can print the "End of program" string afterwards
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of program")

	//Here is an example of using a continue in a for loop:

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// Skip even numbers
			continue
		}
		fmt.Println("Odd number:", i)
	}

	/*
		In this example, the loop iterates from 0 to 9. When i is an even number, the continue statement is executed, which skips the rest of the loop body and continues with the next iteration of the loop. As a result, only odd numbers are printed.
	*/

	/*Here is an example of using a for loop as a for each loop to calculate the total of a shopping cart */

	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)

	result := calculateTotal(cart)

	fmt.Println(result)
}
