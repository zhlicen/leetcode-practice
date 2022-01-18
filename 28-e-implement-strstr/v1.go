package leetcode

func strStr(haystack string, needle string) int {
	var p0, p1 int16
	if len(needle) == 0 {
		return 0
	}
	for p0 = 0; int(p0) < len(haystack); p0++ {
		if haystack[p0] != needle[p1] {
			continue
		}
		for p1 = 1; int(p1) < len(needle); p1++ {
			if int(p0+p1) >= len(haystack) {
				return -1
			}
			if haystack[p0+p1] != needle[p1] {
				p1 = 0
				break
			}
		}
		if p1 != 0 {
			return int(p0)
		}
		continue
	}
	return -1
}
