package main

import "fmt"

func main() {
	str := "foobar"
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
}
