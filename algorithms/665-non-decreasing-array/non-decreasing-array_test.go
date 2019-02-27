package _65_non_decreasing_array

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	{
		[]int{2, 3, 3, 2, 4},
		true,
	},

	{
		[]int{3, 4, 2, 3},
		false,
	},

	{
		[]int{1, 2, 3},
		true,
	},

	{
		[]int{4, 2, 3},
		true,
	},

	{
		[]int{4, 2, 1},
		false,
	},

	// 可以有多个 testcase
}

// 测试是否可以通过
func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, checkPossibility(tc.nums), "输入:%v", tc)
	}
}
