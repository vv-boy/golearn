package strings_stu

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReplace(t *testing.T) {
	str := "abc123abc456"
	assert.Equal(t, strings.Replace(str, "abc", "", 1), "123abc456")
	assert.EqualValues(t, strings.ReplaceAll(str, "abc", ""), "123456")
}