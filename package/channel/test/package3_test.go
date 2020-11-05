package test

import (
	"fmt"
	test2 "stu/package/channel/tutorial2"
	"testing"
)

func TestPackage3(t *testing.T) {
	res := test2.Add(1, 2)
	fmt.Printf("----- %d", res)
}
