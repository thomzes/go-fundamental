package main

import (
	"fmt"
	"go-fundamental/calculation"
	"go-fundamental/condition"
)

func main() {
	fmt.Println("Halo, belajar Golang")

	result := calculation.Add(8, 9)

	fmt.Println(result)

	fmt.Println("Kuis Multiple")

	resultMultiple := calculation.Multiple(5, 5)

	fmt.Println(resultMultiple)

	// IF ELSE CODITION
	resultIf := condition.IfElse(100)
	
	fmt.Println(resultIf)

	// SWITCH CONDITION
	resultSwitch := condition.Switch(1)

	fmt.Println(resultSwitch)

	}

















	
}