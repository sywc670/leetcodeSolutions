package code

// lc 20. 有效的括号
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, r := range s {
		switch r {
		case '{', '(', '[':
			stack = append(stack, r)
		case '}', ']', ')':
			if len(stack) <= 0 || stack[len(stack)-1] != pairs[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
