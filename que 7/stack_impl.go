package main

import (
	"fmt"
	. "strconv"
)

type node struct {
	value string
	next  *node
}

type Stack struct {
	top     *node
	size    int
	maxsize int
}

func (stack *Stack) initialize(maxsize int) {
	stack.maxsize = maxsize
}

func (stack *Stack) push(value string) {
	if stack.size >= stack.maxsize {
		fmt.Println("Stack overflow")
		return
	}
	stack.top = &node{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) pop() (value string) {
	if stack.size > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
	println("underflow")
	return ""
}

func (stack *Stack) String() string {
	str := ""
	size := stack.size
	top := stack.top
	for size > 0 {
		str += " [" + Itoa(size) + " " + stack.top.value + "]"
		size--
		top = top.next
	}
	return str
}

func main() {
	stack := new(Stack)
	stack.initialize(5)
	stack.push("hello")
	stack.push("there")
	fmt.Print("my stack", stack)
}
