package sort_stu

import (
	"fmt"
	"sort"
	"testing"
)

func TestSlice(t *testing.T) {
	age := []int{5, 1, 6, 7, 3, 4}
	// 每次要写 sort.IntSlice 转换类型还是
	sort.Sort(sort.IntSlice(age))
	fmt.Println(age)

	age = []int{5, 1, 6, 7, 3, 4}
	//sort.Ints(age)


	fmt.Println(age)
}
