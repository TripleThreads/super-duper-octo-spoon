package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "asSASA ddd dsjkdsjs dk"

	fmt.Println(len(name))

	count := utf8.RuneCountInString(name)
	fmt.Println(count)

}
