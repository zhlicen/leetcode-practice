package leetcode

func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{1}
	}
	add := false
	digits[l-1]++
	if digits[l-1] >= 10 {
		digits[l-1] -= 10
		add = true
	}
	if l == 1 && add {
		r := []int{1}
		r = append(r, digits...)
		return r
	}
	for i := l - 2; i >= 0; i-- {
		if add {
			digits[i]++
			add = false
		} else {
			return digits
		}
		if digits[i] >= 10 {
			digits[i] -= 10
			add = true
		}
		if i == 0 && add {
			r := []int{1}
			r = append(r, digits...)
			return r
		}
	}
	return digits
}
