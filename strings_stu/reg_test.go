package strings_stu

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegex0(t *testing.T) {

	reg, _ := regexp.Compile("hello")
	fmt.Println(reg.MatchString("hello world"))

	fmt.Println(reg.FindString("hello world hello wei"))

	fmt.Println(regexp.Match("hello", []byte("hello world")))

	fmt.Println(reg.FindIndex([]byte("world hello")))

	fmt.Println(reg.FindAllIndex([]byte("world hello, wei hello"), -1))

}
