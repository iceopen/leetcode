package _16_remove_duplicate_letters

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	s   string
	ans string
}{
	{"bcabc", "abc"},
	{"cbacdcbc", "acdb"},
}

func TestRemoveDuplicateLetters(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		fmt.Println(tc.s, tc.ans)
		ast.Equal(tc.ans, RemoveDuplicateLetters(tc.s), "输入:%v", tc)
	}
}
