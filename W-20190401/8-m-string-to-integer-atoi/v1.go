package leetcode

import "math"

func myAtoi(str string) int {
	var ret int
	var minus, trim, minusCheck bool
	for _, r := range str {
		if !trim && r == ' ' {
			continue
		}
		trim = true
		if !minusCheck {
			minusCheck = true
			if r == '-' {
				minus = true
				continue
			} else if r == '+' {
				continue
			}
		}
		if r > '9' || r < '0' {
			break
		}
		ret = ret*10 + int(r-'0')
		if minus && ret >= -math.MinInt32 {
			ret = -math.MinInt32
			break
		}
		if !minus && ret >= math.MaxInt32 {
			ret = math.MaxInt32
			break
		}
	}
	if minus {
		return -ret
	}
	return ret
}
