package main

import (
	"fmt"
	"go-fundamental/array"
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
	// END IF ELSE CONDITION

	// SWITCH CONDITION
	resultSwitch := condition.Switch(1)

	fmt.Println(resultSwitch)
	// END SWITCH CONDITION

	// FOR LOOPING
	for i := 1; i <= 100; i++ {
		fmt.Println("Looping index ke-", i)
	}
	// END FOR LOOPING

	// WHILE LOOP IN GO (CUSTOM)
	j := 1
	for j <= 100 {
		fmt.Println("While loop ke;", j)
		j++
	}
	// END WHILE LOOP IN GO (CUSTOM)

	// FOREACH/RANGE IN GO
	fmt.Println("FOREACH/RANGE IN GO")
	title := "Golang the best language"

	for index, letter := range title {
		fmt.Println("index :", index, " letter :", string(letter))
	}
	// note = kalo ngga mau pake index, bisa gunain _ (underscore) agar tidak error
	// END FOREACH/RANGE IN GO

	// Kuis Genap Ganjil
	fmt.Println("Kuis Genap Ganjil")

	titleKuis := "Golang the best language"

	for index, letterKuis := range titleKuis {
		if index % 2 == 0 {
			fmt.Println(string(letterKuis))
		} else {
			fmt.Println("ganjil")
		}
	}
	// END KUIS GENAP GANJIL

	// Array in Go
	array.Array()

	// END ARRAY IN GO


	
	









}
