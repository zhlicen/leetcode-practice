package leetcode

import "strconv"

func isSimple(a, b int) bool {
	if a == 1 {
		return true
	}
	for d := 2; d <= a; d++ {
		if a%d == 0 && b%d == 0 {
			return false
		}
	}
	return true
}

func simplifiedFractions(n int) []string {
	result := []string{}
	for i := n; i > 0; i-- {
		for j := 1; j < i; j++ {
			// 判断是否最简，有公约数
			if !isSimple(j, i) {
				continue
			}
			result = append(result, strconv.Itoa(j)+"/"+strconv.Itoa(i))
		}
	}
	return result
}
