package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func main() {

	theAnswer := 42
	var result string
	//You do not need parens surrounding the conditional logic for an if statement:
	if theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)
	//You are also able to initialize a variable inside of an if statement:
	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

	//Now, for switch statements:
	//	rand.Seed(time.Now().Unix())
	//now, using the rand package again, we are using the function Intn() and passing in a ceiling of 7, and adding 1 to it, meaning it will produce a random number between 1
	//and 7
	//	dow := rand.Intn(7) + 1
	//fmt.Println("Day", dow)

	var result2 string

	//you can also initialize a variable inside of the switch statement the same way you can for the if statement:
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result2 = "It's Sunday"
	case 2:
		result2 = "It's Monday"
	case 3:
		result2 = "It's Tuesday"
	case 4:
		result2 = "It's Wednesday"
	case 5:
		result2 = "It's Thursday"
	case 6:
		result2 = "It's Friday"
	case 7:
		result2 = "It's Saturday"
	//	fallthrough
	default:
		result2 = "It is some other day!"
	}
	//There is no use of the break command. Instead, use the fallthrough key word to allow you to skip that case and move to the next.
	fmt.Println(result2)

}
