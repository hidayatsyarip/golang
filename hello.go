package main

import "fmt"

func main() {
	x := [5]float64{
		98,
		76,
		46,
		56,
		72,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
		fmt.Println("value ", value)
	}

	fmt.Println(total / float64(len(x)))
}
