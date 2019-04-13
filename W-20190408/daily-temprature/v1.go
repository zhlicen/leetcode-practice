package leetcode

func dailyTemperatures(T []int) []int {
	l := len(T)
	if l == 0 {
		return []int{}
	}
	for i := 0; i < l; i++ {
		if i == l-1 {
			T[i] = 0
			break
		}
		for j := i + 1; j < l; j++ {
			if T[j] > T[i] {
				T[i] = j - i
				break
			}
			if j == l-1 {
				T[i] = 0
			}
		}
	}
	return T
}
