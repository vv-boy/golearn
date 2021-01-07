package slice_test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := make([]interface{}, 10)
	fmt.Println(cap(arr))

	var mans []string
	man := []string{"vvtest", "vvtest2"}
	fmt.Println(len(man), cap(man))
	mans = man
	fmt.Println(mans)
	mans = append(mans, "vvtest3");

	fmt.Println(man)

}