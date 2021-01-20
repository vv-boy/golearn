package container_test

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var l list.List

	elem := l.PushBack("1231")
	fmt.Print(elem)
}
