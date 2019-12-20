package main

import (
	"strings"
)

type increment func(int) int
type changeCase func(string) string

func double(num int) int {
	return num * 2
}
func upperCase(str string) string {
	return strings.ToUpper(str)
}

func mapFunc(fn increment, array []int) []int {
	newArray := make([]int, len(array))
	for i, val := range array {
		newArray[i] = fn(val)
	}
	return newArray
}

func stringMap(fn changeCase, array []string) []string {
	newArray := make([]string, len(array))
	for i, val := range array {
		newArray[i] = fn(val)
	}
	return newArray
}

func main() {
	res := mapFunc(double, []int{12, 24})
	for _, i := range res {
		println(i)
	}
	resS := stringMap(upperCase, []string{"hello", "there"})

	for _, i := range resS {
		println(i)
	}
}
