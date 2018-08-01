package main

import (
	"fmt"
)

func main() {
	var words = "Hello World!"
	fmt.Printf(words)
	var b = words[0:5]
	var c = words[7:]
	fmt.Println(b)
	fmt.Println(c)

	car := "BMW"
	fmt.Printf(car)
}