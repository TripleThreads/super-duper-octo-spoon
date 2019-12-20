package main

import "fmt"

func ascOrder(num1 int, num2 int) []int {
	if num1 < num2 {
		return []int{num1, num2}
	}
	return []int{num2, num1}
}

func main() {
	fmt.Println(ascOrder(2, 7))
	fmt.Println(ascOrder(7, 2))
}
