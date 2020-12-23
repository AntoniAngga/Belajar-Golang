package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Hallo, Studying Golang")

	result := calculation.Add(5, 10)
	multi := calculation.Multiplication(5, 5)

	fmt.Println(result)
	fmt.Println(multi)
}
