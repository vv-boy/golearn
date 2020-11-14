package strings_stu

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	str := "123abc123"
	// 当长度小于 beteagl.MaxBruteForce 的时候，实际上使用直接对比也是很快的 或者 substr 一个字符串的时候也是
	index := strings.Index(str, "abc")
	fmt.Println(index)

	index = strings.IndexAny(str, "123")
	fmt.Println("any -- " + strconv.Itoa(index))

}
