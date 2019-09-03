package main

import "fmt"

func main() {

	x := [5]float64{
		87,
		82,
		45,
		34,
		21,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

}
