package _976_largest_perimeter_triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

// 测试方法是否可以通过
func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: []int{2, 1, 2},
			},
			a: ans{
				one: 5,
			},
		},
		{
			p: para{
				one: []int{1, 2, 1},
			},
			a: ans{
				one: 0,
			},
		},
		{
			p: para{
				one: []int{3, 2, 3, 4},
			},
			a: ans{
				one: 10,
			},
		},
		{
			p: para{
				one: []int{3, 6, 2, 3},
			},
			a: ans{
				one: 8,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, largestPerimeter(p.one), "输入:%v", p)
	}
}
