package main

import "fmt"

func printAllNumbers(array ...int) {
	for _, i := range array {
		fmt.Println(i)
	}
}

func main() {
	printAllNumbers(1, 2, 4, 5, 6, 7)
}
