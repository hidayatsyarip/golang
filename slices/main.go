package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2) // define a slices, 2 is number of "room for element"
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}
