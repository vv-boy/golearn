package gostr

import (
	"fmt"
	"strings"
	"testing"
)

func TestReplaceStr(t *testing.T) {
	str := strings.Replace("test123", "test", "", 1)
	fmt.Println(str)
}