package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "strings"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	//creating Get response object
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//We will get back a response type pointing to a pointer of *http.Response
	fmt.Printf("Response type: %T", resp)

	defer resp.Body.Close()
	//We are reading the body of the json object that we got using the http.Get()
	//	bytes, err := io.ReadAll(resp.Body)
	//	if err != nil {
	//		panic(err)
	//	}
	//We are converting the content of the body to a string using the string() and printing it to the screen:
	//	content := string(bytes)
	//	fmt.Println(content)

	var tour Tour
	err = toursFromJson(resp.Body, &tour)
	if err != nil {
		panic(err)
	}

	fmt.Println("Title:", tour.Title)
	fmt.Println("Body:", tour.Body)

}

func toursFromJson(r io.Reader, tour *Tour) error {
	decoder := json.NewDecoder(r)

	return decoder.Decode(tour)
}

// declaring a new type of Tour, and two fields: Name and Price.  Because they are both strings, I can declare them on the same line:
type Tour struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
