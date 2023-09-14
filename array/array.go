package array

import "fmt"

func Array() {
	languages := [...]string{
		"Ruby",
		"Python",
		"Javascript",
		"Go",
		"C#",
		"Kotlin",
	}

	for index, lang := range languages {
		fmt.Println("Index ke :", index, "Language :",string(lang))
	}

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println("Panjang index :", length)

}