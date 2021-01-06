package gostr

import (
	"fmt"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	str := "abcdefg"
	reader := strings.NewReader(str)
	fmt.Println(reader.Len())
	fmt.Println(str)
}
