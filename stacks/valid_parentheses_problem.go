package stacks

func isValid(s string) bool {
	stack := Stack{}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack.push(int(char))
		} else {
			if stack.isEmpty() {
				return false
			}

			top := stack.peek()

			if (char == ')' && top == '(') ||
				(char == '}' && top == '{') ||
				(char == ']' && top == '[') {
				stack.pop()
			} else {
				return false
			}
		}
	}

	return stack.isEmpty()
}
