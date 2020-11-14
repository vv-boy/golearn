package strings_stu

import (
	"fmt"
	"strings"
	"testing"
)

func TestReplacer(t *testing.T) {
	str := "abc123abc"
	replacer := strings.NewReplacer("abc", "456")
	fmt.Println(replacer.Replace(str))


}