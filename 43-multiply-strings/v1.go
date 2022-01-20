package leetcode

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	var muls = make([][]byte, l1*l2)
	idx := 0
	// 计算出所有的相乘，然后合并
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			r := (num2[j] - '0') * (num1[i] - '0')
			reserveLen := (l1 - 1 - i) + (l2 - 1 - j)
			// mul 存储逆序的按位乘法结果
			mul := make([]byte, 2, 2)
			// 最多两位数
			// 第一个byte放后边0的个数
			// 第二个byte放两数乘积
			mul[0] = byte(reserveLen)
			mul[1] = r
			muls[idx] = mul
			idx++
		}
	}
	var mulIdx byte = 0
	last := false
	// 最大长度为l1+l2
	maxResultLen := l1 + l2
	var result []byte = make([]byte, maxResultLen, maxResultLen)
	// 进位
	var moreAdd int = 0
	// 单次求和
	var singleSum int = 0
	for !last && maxResultLen > 0 {
		// 默认为最后一位
		last = true
		for _, mul := range muls {
			// 还原数字
			// 前边0的数量
			mulLen := mul[0] + 1
			if mul[1] >= 10 {
				mulLen++
			}
			if mulLen > mulIdx {
				last = false
				if mulIdx >= mul[0] {
					numIdx := mulIdx - mul[0]
					if numIdx == 0 {
						singleSum += int(mul[1] % 10)
					} else if numIdx == 1 {
						singleSum += int(mul[1] / 10)
					}
				}
			}
		}
		// 加进位
		singleSum += moreAdd
		// 取新进位
		moreAdd = int(singleSum / 10)
		// 推入结果
		result[maxResultLen-1] = byte(singleSum%10) + '0'
		maxResultLen--
		singleSum = 0
		mulIdx++
	}
	// trim left 0
	maxResultLen = len(result)
	for i := 0; i < maxResultLen; i++ {
		if result[i] != '0' && result[i] != 0 {
			result = result[i:]
			break
		}
	}

	return string(result)
}
