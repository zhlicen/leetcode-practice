package leecode

import "fmt"

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	l := len(A)
	s := false
	fc := []byte{}
	m := make(map[byte]interface{})
	for i := 0; i < l; i++ {
		if !s {
			if _, ok := m[A[i]]; ok {
				fmt.Println("sssss")
				s = true
			} else {
				m[A[i]] = nil
			}
		}

		if A[i] != B[i] {
			if len(fc) > 4 {
				return false
			}
			fc = append(fc, A[i], B[i])
		}
	}
	fmt.Println(fc)
	if len(fc) > 4 {
		return false
	}
	if len(fc) == 4 {
		if fc[0] == fc[3] && fc[1] == fc[2] {
			return true
		} else {
			return false
		}
	}

	return s
}
