package main

import "fmt"

func main() {
	a := 1
	goto TARGET

TARGET:
	b := 9

	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
