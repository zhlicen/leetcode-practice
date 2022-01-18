package leetcode

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	f := make(map[int]interface{})
	for i, n := range nums {
		if _, ok := f[i]; ok {
			continue
		}
		if i == len(nums)-1 {
			return nums[i]
		}
		for j := i + 1; j < len(nums); j++ {

			if n == nums[j] {
				f[j] = nil
				break
			}
			if j == len(nums)-1 {
				return n
			}
		}
	}
	return -1
}
