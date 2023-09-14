package kuis

import "fmt"

func GenapGanjil(title string) string {
	fmt.Println("Kuis Genap Ganjil")

	for index, letter := range title {
		if index % 1 == 0 {
			return string(letter)
		} else {
			return "test"
		}
	}
	return "3"
}