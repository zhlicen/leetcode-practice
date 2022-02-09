package leetcode

func reverse(x int) int {
	minus := x < 0
	if minus {
		x = -x
	}
	maxInt32 := 1<<31 - 1
	maxInt32By10 := maxInt32 / 10
	var result int
	for x != 0 {
		unit := x % 10
		x = x / 10
		if result > maxInt32By10 {
			return 0
		} else if result == maxInt32By10 {
			// 判断临界值
			if !minus {
				if unit > maxInt32%10 {
					return 0
				}
			} else {
				if unit > maxInt32%10+1 {
					return 0
				}
			}

		}
		result = result*10 + unit
	}

	if minus {
		result = -result
	}
	return int(result)
}
