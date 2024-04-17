package main

import (
	"fmt"
	"sort"
)

func main() {

	//Maps are unordered collections of key value pairs.  It is a Hash table that lets you store collections of data and then arbitrarily
	//find items in the collection using their keys.

	//It is common to use strings for keys, and any other type for the associated values, but keys can be of any type

	states := make(map[string]string)

	states["WA"] = "Washington"
	states["TN"] = "Tennessee"
	states["AL"] = "Alabama"
	fmt.Println(states)

	//You can also assign variables to keys in a map

	tennessee := states["TN"]

	fmt.Println(tennessee)

	//If you want to remove an item from a map, use the builtin delete()

	delete(states, "TN")

	//you can also add to a map like so

	states["NY"] = "New York"

	fmt.Println(states)

	//if you want to iterate through a map, you can use a for loop where k is the key and v is the value

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	//if you are wanting to display the values of a map in a certain order, you will have to do it yourself:
	//create a slice of an array using the make() and the initial value as the length of the states map
	//initialize  a variable i for the current index of the slice and set it to 0
	//use a for loop and have k := range states, and set the index of keys in the loop to [i], and set it to k (the current key in the states map)
	//continue through the loop till you get to the end of the map using i++
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println(keys)
	//if you are wanting to sort the values in a certain order, you have to use the sort package
	sort.Strings(keys)
	fmt.Println(keys)
	//Now iterate through the slice using i for the index, and print these out to the console.
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
