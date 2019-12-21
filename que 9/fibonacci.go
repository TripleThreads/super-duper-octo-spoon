package main

import "fmt"

func fibonacci(num int) {
	prev := 1
	next := 1
	current := 0
	fmt.Printf("%d %d ", prev, next)
	for i := 1; i < num; i++ {
		current = next
		next = next + prev
		prev = current
		fmt.Printf("%d ", next)
	}
}

func main() {
	fibonacci(10)
}
