package leetcode

func isValid(s string) bool {
	idx := -1
	maxArrLen := len(s)/2 + 1
	if maxArrLen == 0 {
		return true
	}
	arr := make([]byte, maxArrLen)
	l := len(s)
	for i := 0; i < l; i++ {
		switch s[i] {
		case '[':
			idx++
			if idx >= maxArrLen {
				return false
			}
			arr[idx] = ']'
		case '{':
			idx++
			if idx >= maxArrLen {
				return false
			}
			arr[idx] = '}'
		case '(':
			idx++
			if idx >= maxArrLen {
				return false
			}
			arr[idx] = ')'
		case ']', '}', ')':
			if idx == -1 {
				return false
			}
			if arr[idx] == s[i] {
				idx--
			} else {
				return false
			}
		}

	}
	if idx == -1 {
		return true
	}
	return false
}
