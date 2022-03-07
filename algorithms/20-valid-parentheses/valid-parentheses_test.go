package valid_parentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one bool
}

// Test_IsOK 测试是否通过
func Test_IsOK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{"()[]{}"},
			ans{true},
		},
		{
			para{"(]"},
			ans{false},
		},
		{
			para{"({[]})"},
			ans{true},
		},
		{
			para{"(){[({[]})]}"},
			ans{true},
		},
		{
			para{"((([[[{{{"},
			ans{false},
		},
		{
			para{"(())]]"},
			ans{false},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, isValid(p.one), "输入:%v", p)
	}
}
