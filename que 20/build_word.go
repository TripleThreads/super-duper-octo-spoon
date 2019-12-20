package main

import (
	"strings"
)

func buildWord(word string, array []string) int {

	stack := new(Stack)
	stack.initialize(len(array))
	index := 0
	var removedWords []string

	for index < len(word) {
		pushed := false

		for _, segment := range array {

			itsRemoved := false

			for _, w := range removedWords {
				if w == segment {
					itsRemoved = true
					break
				}
			}
			if strings.Index(word, segment) == index && !itsRemoved {
				stack.Push(segment)
				index += len(segment)
				pushed = true
			}
		}
		if !pushed && stack.Size > 0 {
			removed := stack.Pop()
			index -= len(removed)
			removedWords = append(removedWords, removed)
		}
		if !pushed && stack.Size == 0 {
			return 0
		}
	}
	return stack.Size
}
func main() {
	index := buildWord("buildword", []string{"buil", "dwor", "bu", "ild", "wo", "rd"})
	println(index)
}
