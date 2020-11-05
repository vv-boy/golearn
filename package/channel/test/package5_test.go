package test

import (
	"fmt"
	. "stu/package/channel/tutorial"
	"testing"
)

func TestPackage5(t *testing.T) {
	res := Add(1, 2)
	fmt.Printf("---- %d", res)
}
