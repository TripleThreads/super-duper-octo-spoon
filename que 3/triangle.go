package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Print("\n")
	}
}
