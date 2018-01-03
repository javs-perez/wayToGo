package main

import "fmt"

func main() {
	str := "Go is a beautiful languages!"
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}

	fmt.Println()
	str2 := "Chinese: 汉字"
	for pos, char := range str2 {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
	fmt.Println()
	fmt.Println("index\tint(rune)\t\t\trune\t\tchar\t\tbytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d\t\t%d\t\t\t\t\t%U\t\t'%c'\t\t% X \n", index, rune, rune, rune, []byte(string(rune)))
	}
}
