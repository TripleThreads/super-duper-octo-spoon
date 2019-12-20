package main

import "fmt"

func plusTwo() func(int) int {
	return func(x int) int {
		return x + 2
	}
}

func plusX() func(int) int {
	return func(x int) int {
		return x + x
	}
}
func main() {
	p := plusTwo()
	x := plusX()
	fmt.Println(p(2))
	fmt.Println(x(3))
}
