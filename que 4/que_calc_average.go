package main

import "fmt"

func main() {
	array := []float64{101, 3, 77, 82, 12}

	total := 0.0
	for _, item := range array {
		total += item
	}
	fmt.Println(total / float64(len(array)))
}
