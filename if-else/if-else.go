package main

import "fmt"

func main() {
	a := 7.0
	b := 9.0

	fmt.Println(a, b)

	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("b > a")
	}
}
