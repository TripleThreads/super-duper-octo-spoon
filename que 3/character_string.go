package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "asSASA ddd dsjkdsjs dk"

	dict := make(map[int32]int)

	for _, char := range name {
		dict[char]++
	}

	for key, value := range dict {
		fmt.Printf("character %c appeared %d times \n", key, value)
	}
	fmt.Println(len(name))

	count := utf8.RuneCountInString(name)
	fmt.Println(count)

}
