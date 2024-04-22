package main

import (
	"fmt"
)

func main() {

	//In Go, you can make a multiline string using a ` and a ` to end:

	//	aSentence := `
	//		This is the best thing ever.
	//		Wow, I cannot believe I can do this.
	//		Amazing
	//	`

	count := 0
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			n := a * b

			s := fmt.Sprintf("%d", n)

			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Println(count)


		


}
