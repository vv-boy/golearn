package strings_stu

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStrconv(t *testing.T) {
	str := "abc1231"
	num, _ := strconv.Atoi(str)
	fmt.Println(num)

	num2, err := strconv.ParseInt("128", 10, 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num2)

	s := "128"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	fmt.Println(strconv.FormatInt(128, 16))
	fmt.Println(strconv.FormatInt(15, 8))

	fmt.Println(strconv.Itoa(18));

}
