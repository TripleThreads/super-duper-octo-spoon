package stack

// this file is here because i needed to import in several packages
import (
	"fmt"
	. "strconv"
)

type Node struct {
	value string
	next  *Node
}

type Stack struct {
	Top     *Node
	Size    int
	MaxSize int
}

func (stack *Stack) Initialize(maxsize int) {
	stack.MaxSize = maxsize
}

func (stack *Stack) Push(value string) {
	if stack.Size >= stack.MaxSize {
		fmt.Println("Stack overflow")
		return
	}
	stack.Top = &Node{
		value: value,
		next:  stack.Top,
	}
	stack.Size++
}

func (stack *Stack) Pop() (value string) {
	if stack.Size > 0 {
		value = stack.Top.value
		stack.Top = stack.Top.next
		stack.Size--
		return
	}
	println("underflow")
	return ""
}

func (stack *Stack) String() string {
	str := ""
	size := stack.Size
	top := stack.Top
	for size > 0 {
		str += " [" + Itoa(size) + " " + stack.Top.value + "]"
		size--
		top = top.next
	}
	return str
}
