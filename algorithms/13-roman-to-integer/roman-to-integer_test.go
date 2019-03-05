package romanToInteger

import (
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
	one int
}

// Test_IsOK 测试是否通过
func Test_IsOK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{"XXXIX"},
			ans{39},
		},
		{
			para{"MDCCCLXXXVIII"},
			ans{1888},
		},
		{
			para{"MCMLXXVI"},
			ans{1976},
		},
		{
			para{"MMMCMXCIX"},
			ans{3999},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, romanToInt(p.one), "输入:%v", p)
	}
}
