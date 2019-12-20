package main

import "fmt"

func main() {
	i := 0

label:
	fmt.Println(i)

	if i < 10 {
		i++
		goto label
	}
}
