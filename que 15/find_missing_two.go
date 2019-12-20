package main

import "fmt"

func main() {
	array := []int{4, 2, 3}

	count := make([]int, len(array)+2)

	for _, i := range array {
		count[i-1]++
	}

	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			fmt.Println(i + 1)
		}
	}
}
