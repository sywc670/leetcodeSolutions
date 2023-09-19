package code

import "strings"

// lc 151
func reverseWords(s string) string {
	strSlice := strings.Fields(s)
	l := len(strSlice)
	for i := 0; i < l/2; i++ {
		strSlice[i], strSlice[l-i-1] = strSlice[l-i-1], strSlice[i]
	}
	return strings.Join(strSlice, " ")
}

func reverseWordsV1(s string) string {
	var builder strings.Builder
	strSlice := strings.Fields(s)
	for l := len(strSlice) - 1; l > 0; l-- {
		builder.WriteString(strSlice[l])
		builder.WriteString(" ")
	}
	builder.WriteString(strSlice[0])
	return builder.String()
}

// lc 71. 简化路径
func simplifyPath(path string) string {
	stack := []string{}
	for _, name := range strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	}) {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			if name != "." && name != "" {
				stack = append(stack, name)
			}
		}
	}
	return "/" + strings.Join(stack, "/")
}

// lc 1768. 交替合并字符串
func mergeAlternately(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	newWord := make([]byte, 0, m+n)
	p1, p2 := 0, 0
	l := min(m, n)
	for i := 0; i < 2*l; i++ {
		if i%2 == 0 {
			newWord = append(newWord, word1[p1])
			p1++
		} else {
			newWord = append(newWord, word2[p2])
			p2++
		}
	}
	for p1 < m {
		newWord = append(newWord, word1[p1])
		p1++
	}
	for p2 < n {
		newWord = append(newWord, word2[p2])
		p2++
	}
	return string(newWord)
}

// optimize
func mergeAlternatelyV2(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	newWord := make([]byte, 0, m+n)
	for i := 0; i < m || i < n; i++ {
		if i < m {
			newWord = append(newWord, word1[i])
		}
		if i < n {
			newWord = append(newWord, word2[i])
		}
	}
	return string(newWord)
}
