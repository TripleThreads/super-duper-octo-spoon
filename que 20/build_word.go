package main

import (
	stack "awesomeProject"
	"strings"
)

func buildWord(word string, array []string) int {

	wordStack := new(stack.Stack)
	wordStack.Initialize(len(array))
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
				wordStack.Push(segment)
				index += len(segment)
				pushed = true
			}
		}
		if !pushed && wordStack.Size > 0 {
			removed := wordStack.Pop()
			index -= len(removed)
			removedWords = append(removedWords, removed)

		} else if !pushed && wordStack.Size == 0 {
			return 0
		}
	}
	return wordStack.Size
}
func main() {
	count := buildWord("buildword", []string{"buil", "dwor", "bu", "ild", "wo", "rd"})
	println(count)
}
