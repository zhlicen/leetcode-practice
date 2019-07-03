package leetcode

import (
	"fmt"
	"testing"
)

func Test_exist(t *testing.T) {
	// input1 := [][]byte{
	// 	{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	// }
	// input2 := 'ABCCE'
	input1 := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	input2 := "ABCESEEEFS"
	// input1 := [][]byte{
	// 	{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'},
	// }
	// input2 := 'AAB'
	fmt.Println(exist(input1, input2))
}
