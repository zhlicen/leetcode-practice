package leetcode

func buddyStrings_v2(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	l := len(A)
	s := false
	fc := make([]byte, 4)
	fcl := 0
	m := make(map[byte]interface{})
	for i := 0; i < l; i++ {
		if !s {
			if _, ok := m[A[i]]; ok {
				s = true
			} else {
				m[A[i]] = nil
			}
		}

		if A[i] != B[i] {
			if fcl >= 4 {
				return false
			}
			fc[fcl] = A[i]
			fc[fcl+1] = B[i]
			fcl += 2
		}
	}
	if fcl == 4 {
		if fc[0] == fc[3] && fc[1] == fc[2] {
			return true
		} else {
			return false
		}
	}

	return s
}
