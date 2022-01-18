package leetcode

import "strings"

func toGoatLatin(S string) string {
	ss := strings.Split(S, " ")
	for i, s := range ss {
		if s == "" {
			continue
		}
		if s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' || s[0] == 'A' || s[0] == 'E' || s[0] == 'I' || s[0] == 'O' || s[0] == 'U' {
			ss[i] = s + "ma"
		} else {
			ss[i] = s[1:] + s[0:1] + "ma"
		}
		for j := 0; j < i+1; j++ {
			ss[i] += "a"
		}
	}
	return strings.Join(ss, " ")
}
