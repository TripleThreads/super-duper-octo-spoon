package main

func main() {
	reverse := ""
	subRev := ""

	stack := new(Stack)
	input := "foo(bar(baz))blim"

	stack.initialize(len(input))

	for _, char := range input {

		if stack.Size == 0 && string(char) != "(" {
			reverse += string(char)
		}

		if string(char) == "(" || (string(char) != ")" && stack.Size > 0) {
			stack.Push(string(char))
		}

		if string(char) == ")" {
			c := stack.Pop()
			for c != "(" {
				subRev += c
				c = stack.Pop()
			}
			if stack.Size == 0 {
				reverse += subRev
				continue
			}
			for _, c := range subRev {
				stack.Push(string(c))
			}
			subRev = ""
		}
	}

	println(reverse)
}
