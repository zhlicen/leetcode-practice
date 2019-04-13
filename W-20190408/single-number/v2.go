package leetcode

func singleNumber(nums []int) int {
	l := len(nums)
	f := make([]bool, l)
	for i, n := range nums {
		if f[i] {
			continue
		}
		for j := i + 1; j < l; j++ {
			if n == nums[j] {
				f[j] = true
				break
			}
			if j == l-1 {
				return n
			}
		}
	}
	return nums[l-1]
}
