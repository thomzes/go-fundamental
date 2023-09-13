package condition

import "fmt"

func IfElse(score int) string {
	// IF ELSE CONDITION
	fmt.Println("IF ELSE CONDITION")
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	return grade
}
