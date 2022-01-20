package leetcode

func multiplyV2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// 短的放下边
	if len(num1) < len(num2) {
		n := num1
		num1 = num2
		num2 = n
	}
	l1 := len(num1)
	l2 := len(num2)
	var rows = make([][]byte, 0, l2)
	maxResultLen := 0
	// num2 按位乘以 num1
	for i := l2 - 1; i >= 0; i-- {
		// 进位
		var moreAdd byte = 0
		row := make([]byte, l2-i-1, l2-i-1+l1+1)
		for j := l1 - 1; j >= 0; j-- {
			r := (num2[i] - '0') * (num1[j] - '0')
			r += byte(moreAdd)
			// 推入结果
			row = append(row, r%10)
			moreAdd = r / 10
			// idx++
		}
		if moreAdd > 0 {
			row = append(row, moreAdd)
		}
		if len(row) > maxResultLen {
			maxResultLen = len(row)
		}
		rows = append(rows, row)
	}
	maxResultLen++
	var colIdx byte = 0
	last := false
	var result []byte = make([]byte, maxResultLen, maxResultLen)
	// 进位
	var moreAdd int = 0
	// 单次求和
	var colResult int = 0
	for !last && maxResultLen > 0 {
		// 默认为最后一位
		last = true
		for _, row := range rows {
			if len(row) > int(colIdx) {
				last = false
				colResult += int(row[colIdx])
			}
		}
		// 加进位
		colResult += moreAdd
		// 取新进位
		moreAdd = int(colResult / 10)
		// 推入结果
		result[maxResultLen-1] = byte(colResult%10) + '0'
		maxResultLen--
		colResult = 0
		colIdx++
	}

	// trim left 0
	for i := 0; i < len(result); i++ {
		if result[i] != '0' && result[i] != 0 {
			return string(result[i:])
		}
	}

	return string(result)
}
