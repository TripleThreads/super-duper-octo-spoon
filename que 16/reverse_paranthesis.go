package main

import stack "awesomeProject"

func main() {
	reverse := ""
	subRev := ""

	_stack := new(stack.Stack)
	input := "foo(bar(baz))blim"

	_stack.Initialize(len(input))

	for _, char := range input {

		if _stack.Size == 0 && string(char) != "(" {
			reverse += string(char)
		}

		if string(char) == "(" || (string(char) != ")" && _stack.Size > 0) {
			_stack.Push(string(char))
		}

		if string(char) == ")" {
			c := _stack.Pop()
			for c != "(" {
				subRev += c
				c = _stack.Pop()
			}
			if _stack.Size == 0 {
				reverse += subRev
				continue
			}
			for _, c := range subRev {
				_stack.Push(string(c))
			}
			subRev = ""
		}
	}

	println(reverse)
}
