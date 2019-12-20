package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {

		if i%15 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
