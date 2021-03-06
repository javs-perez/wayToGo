package main

import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit addres-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { //dereferencig *a to get back to the array is not necessary!
		sum += v
	}
	return
}
