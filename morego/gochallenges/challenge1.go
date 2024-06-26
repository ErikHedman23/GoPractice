package main

import (
	"fmt"
	"strings"
)

func SeparateWords() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	words := strings.Fields(text)

	counts := map[string]int{}

	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}
