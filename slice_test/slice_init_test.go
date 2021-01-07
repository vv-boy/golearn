package slice

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	var obj []string
	obj1 := []string {"test"}
	obj = obj1
	fmt.Print(obj)
}
