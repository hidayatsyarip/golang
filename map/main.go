package main

import "fmt"

func main() {
	x := map[string]string{
		"H": "Honda",
		"S": "Suzuki",
		"T": "Toyota",
		"M": "Mazda",
		"B": "BMW",
	}

	if name, ok := x["B"]; ok {
		fmt.Println(name, ok)
	}
}
