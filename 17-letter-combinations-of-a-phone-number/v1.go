package leetcode

func letterCombinations(digits string) []string {
	code := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	input := [][]byte{}
	for _, num := range digits {
		if num-'0' >= 2 {
			input = append(input, code[num-'0'-2])
		}
	}
	length := len(input)
	if len(input) == 0 {
		return []string{}
	}
	index := make([]int, length)
	resultCount := 1
	for i := 0; i < len(input); i++ {
		resultCount *= len(input[i])
	}
	results := make([]string, resultCount)
	for i := 0; i < resultCount; i++ {
		result := make([]byte, len(input))
		for j := 0; j < length; j++ {
			result[j] = input[j][index[j]]
		}
		results[i] = string(result)
		index[0]++
		for k := 0; k < length; k++ {
			if index[k] < len(input[k]) {
				break
			}
			if k == length-1 {
				return results
			}
			index[k] = 0
			index[k+1]++
			if index[length-1] > len(input[length-1])-1 {
				return results
			}
		}
	}
	return results
}
