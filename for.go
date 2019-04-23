package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}

	b := 0
	for true {
		fmt.Println(b)
		b++
		if b >= 10 {
			break
		}
	}
}
