package leecode

// https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	c := 0
	startIdx := l - 1
	for i := startIdx; i >= 0; i-- {
		if s[i] != ' ' {
			startIdx = i
			break
		}
	}
	for i := startIdx; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		c++
	}
	return c
}
