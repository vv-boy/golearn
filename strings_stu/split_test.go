package strings_stu

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	str := "weiwei,21,man,it"

	fmt.Println(strings.Split(str, ","))

	fmt.Println(strings.SplitN(str, ",", 3))

	fmt.Println(strings.SplitAfter(str, ","))

	fmt.Println(strings.SplitAfterN(str, ",", 2))

}
