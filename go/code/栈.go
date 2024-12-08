package code

import "math"

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

// 155. 最小栈
// 思路1：辅助栈用来保存每次入栈的最小值
// 思路2：栈中保存入栈的值以及最小值
type MinStack []intPair
type intPair [2]int

func ConstructorMinStack() MinStack {
	return MinStack{intPair{-1, math.MaxInt}}
}

func (ms *MinStack) Push(val int) {
	*ms = append(*ms, intPair{val, min(ms.GetMin(), val)})
}

func (ms *MinStack) Pop() {
	*ms = (*ms)[:len(*ms)-1]
}

func (ms MinStack) Top() int {
	return ms[len(ms)-1][0]
}

func (ms MinStack) GetMin() int {
	return ms[len(ms)-1][1]
}
