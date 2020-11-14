package strings_stu

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	str := " abc123abc   "
	str2 := " abc123abc "

	assert.Equal(t, strings.Trim(str, " "), "abc123abc")

	assert.Equal(t, strings.TrimLeft(str2, " "), "abc123abc ")
}
