package main

import "fmt"

func average(slice []float64) float64 {
	total := 0.0
	for _, item := range slice {
		total += item
	}
	return total / float64(len(slice))
}

func main() {
	array := []float64{101, 3, 77, 82, 12}
	fmt.Println(average(array))
}
