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

// 2390. 从字符串中移除星号
func removeStars(s string) string {
	stack := make([]byte, 0, len(s))
	for _, r := range s {
		if r == '*' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
		if r != '*' {
			stack = append(stack, byte(r))
		}
	}
	return string(stack)
}

// 735. 小行星碰撞
func asteroidCollision(asteroids []int) (stack []int) {
	for _, a := range asteroids {
		if a > 0 {
			stack = append(stack, a)
			continue
		}

	Loop:
		for {
			top := len(stack) - 1
			if len(stack) == 0 || stack[top] < 0 {
				stack = append(stack, a)
				break
			}

			if -a == stack[top] {
				stack = stack[:top]
				break Loop
			}

			if -a < stack[top] {
				break Loop
			}

			stack = stack[:top]
		}
	}

	return
}
