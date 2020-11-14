package strings_stu

import (
	"fmt"
	"testing"
	"unicode"
)

func TestUnicode(t *testing.T) {

	fmt.Println(unicode.IsNumber(1))

	fmt.Println(unicode.IsNumber('1'))

	fmt.Println(unicode.IsNumber(49))

	fmt.Println(unicode.IsNumber('\u0049'))
}

