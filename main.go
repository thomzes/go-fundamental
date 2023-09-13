package main

import (
	"fmt"
	"go-fundamental/calculation"
)

func main() {
	fmt.Println("Halo, belajar Golang")

	result := calculation.Add(8, 9)

	fmt.Println(result)

	fmt.Println("Kuis Multiple")

	resultMultiple := calculation.Multiple(5, 5)

	fmt.Println(resultMultiple)

}
