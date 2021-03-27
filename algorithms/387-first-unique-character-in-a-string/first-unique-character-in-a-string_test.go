package _87_first_unique_character_in_a_string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{

	{
		"leetcode",
		0,
	},

	{
		"loveleetcode",
		2,
	},

	{"aabbcc", -1},

	// 可以有多个 testcase
}

func Test_firstUniqChar(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("%v\n", tc)
		ast.Equal(tc.ans, firstUniqChar(tc.s), "输入:%v", tc)
	}
}
