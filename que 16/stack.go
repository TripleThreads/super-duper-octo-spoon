package main

import (
	"fmt"
)

type node struct {
	Value string
	Next  *node
}

type Stack struct {
	Top     *node
	Size    int
	maxSize int
}

func (stack *Stack) initialize(maxSize int) {
	stack.maxSize = maxSize
}

func (stack *Stack) Push(value string) {
	if stack.Size >= stack.maxSize {
		fmt.Println("Stack overflow")
		return
	}
	stack.Top = &node{
		Value: value,
		Next:  stack.Top,
	}
	stack.Size++
}

func (stack *Stack) Pop() (value string) {
	if stack.Size > 0 {
		value = stack.Top.Value
		stack.Top = stack.Top.Next
		stack.Size--
		return
	}
	println("underflow")
	return ""
}
