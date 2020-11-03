package optional_params

import (
	"fmt"
	"reflect"
	"testing"
)

func add(a ...int) {
	fmt.Print(reflect.TypeOf(a))
}
// 得到的是 []int 类型。
func TestAdd(t *testing.T) {
	add(1, 2)
}

