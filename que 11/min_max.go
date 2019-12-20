package main

func findMax(array []int) int {
	max := -1
	for _, i := range array {
		if i > max {
			max = i
		}
	}
	return max
}

func findMin(array []int) int {
	min := 1100000000000
	for _, i := range array {
		if i < min {
			min = i
		}
	}
	return min
}

func main() {

}
