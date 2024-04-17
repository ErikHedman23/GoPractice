package main

import (
	"fmt"
	"sort"
)

func main() {

	var colors = [3]string{"Green", "Red", "Blue"}
	fmt.Println(colors)
	//By removing the length of the array (the [3]), it becomes a slice, and you can add and remove items from it during runtime
	var names = []string{"Bob", "Jeff", "Matthew"}

	names = append(names, "Jeffery")

	fmt.Println(names)
	//to remove an item in a slice, you will also use to append(), but like so:

	names = append(names[1:len(names)])

	fmt.Println(names)

	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 135
	numbers[2] = 136
	numbers[3] = 137
	numbers[4] = 138

	fmt.Println(numbers)

	//Using them make function allows you to input the type of the slice, the original starting amount, and the limit.  If you leave the limit arugument empty, you can add as
	//much as you would like to to the slice by using the append()

	sort.Ints(numbers)
	fmt.Println(numbers)
}
