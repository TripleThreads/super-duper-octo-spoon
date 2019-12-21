package main

import (
	stack "awesomeProject"
	"fmt"
)

func main() {
	demoStack := new(stack.Stack)
	demoStack.Initialize(5)
	demoStack.Push("hello")
	demoStack.Push("there")
	fmt.Print("my demoStack", demoStack)
}
