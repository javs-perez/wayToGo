package main

import (
	"fmt"
)

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	var n int

	Multiply(10, 5, &n)
	fmt.Println("Multiply:", n)
}


