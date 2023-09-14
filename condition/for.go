package condition

import "fmt"

func For(i int) string {
	fmt.Println("FOR LOOP CONDITION")

	for i := 1; i <= 100; i++ {
		return "Saya sedang belajar go, looping ke-"
	}

	return "a"
}
