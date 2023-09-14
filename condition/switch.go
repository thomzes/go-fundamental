package condition

import "fmt"

func Switch(number int) string {
	fmt.Println("SWITCH CONDITION")

	switch number {
	case 1:
		return "Satu"
	case 2:
		return "Dua"
	case 3:
		return "Tiga"
	default:
		return "Deafult"
	}
}
