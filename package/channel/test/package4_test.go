package test

import (
	"fmt"
	mytest "stu/package/channel/tutorial"
	"testing"
)

func TestPackage4(t *testing.T) {
	res := mytest.Add(1, 2)
	fmt.Printf("------ %d", res)
}
