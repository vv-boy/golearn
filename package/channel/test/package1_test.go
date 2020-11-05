package test

import (
	"fmt"
	"testing"
)
import "stu/package/channel/tutorial"
func TestAdd(t *testing.T) {
	res :=test.Add(1, 2)
	fmt.Printf("%d", res)
}