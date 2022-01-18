package leetcode

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	m := make(map[rune]string)
	for _, c := range s {
		if cc, ok := m[c]; ok {
			m[c] = cc + string(c)
		} else {
			m[c] = string(c)
		}
	}
	ss := make([]string, len(m))
	count := 0
	for _, v := range m {
		ss[count] = v
		count++
	}
	sort.SliceStable(ss, func(i, j int) bool {
		return len(ss[i]) > len(ss[j])
	})
	return strings.Join(ss, "")
}
