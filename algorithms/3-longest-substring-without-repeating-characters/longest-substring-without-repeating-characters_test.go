package longest_substring_without_repeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

// Test_OK 测试是否可以通过
func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: "abcabcbb",
			},
			a: ans{
				one: 3,
			},
		},
		{
			p: para{
				one: "bbbbbbbb",
			},
			a: ans{
				one: 1,
			},
		},
		{
			p: para{
				one: "pwwkew",
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, lengthOfLongestSubstring(p.one), "输入:%v", p)
	}
}
