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
