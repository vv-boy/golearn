package strings_stu

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrefix(t *testing.T) {
	str := "abcdefg123"

	fmt.Println(strings.HasPrefix(str, "abc"))

	fmt.Println(strings.HasSuffix(str, "abc"))

	fmt.Println(strings.HasSuffix(str, "123"))

	fmt.Println(strings.HasSuffix(str, "1234"))


}
