package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()

	fmt.Println("I performed this computation at: ", n)

	t := time.Date(2024, time.April, 15, 19, 0, 0, 0, time.UTC)
	fmt.Println("The time of today is ", t)

	//This is printing out a formated version of the one above this.
	fmt.Println(t.Format(time.ANSIC))

	//time.Parse handles errors, so we put in a _ to simply handle any errors, but we are ignoring the error:

	parsedTime, _ := time.Parse(time.ANSIC, "Mon April 15 19:00:00 2024")

	fmt.Printf("The time of parseTime is %T\n", parsedTime)
}
