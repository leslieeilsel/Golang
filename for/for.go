package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for a := 0; a < 10000000000; a++ {
		//fmt.Println(a)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
	b := 0
	for true {
		//fmt.Println(b)
		b++
		if b >= 10000000000 {
			break
		}
	}
}
