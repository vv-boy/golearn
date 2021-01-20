package slice

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	arr := []int{1, 2, 3, 4}

	arr1 := arr[0:2]
	fmt.Print(len(arr1))

	fmt.Print(cap(arr1))
}
