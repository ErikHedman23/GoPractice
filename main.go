package main

import (
	"fmt"
	"math"
)

func main() {
	//if you want to initialize several variables of the same type, you can do so like this:
	i1, i2, i3 := 12, 45, 68

	intSum := i1 + i2 + i3

	fmt.Println("Integer sum =", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3

	floatSum := f1 + f2 + f3

	fmt.Println("Float sum =", floatSum)

	floatSum = math.Round(floatSum*100) / 100

	fmt.Println("The float sum is now =", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	//using a constant in the math package called math.Pi, and we are using Printf() to format the circumference variable using %.2f, which means
	//a floating number with 2 digits after the decimal
	fmt.Printf("Circumference = %.2f\n", circumference)
}
