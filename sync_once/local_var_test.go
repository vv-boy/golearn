package sync_once

import (
	"fmt"
	"sync"
	"testing"
)

func TestLocalVar(t *testing.T) {
	getOnce()
	// 这样写是无效的，因为one 已经重新初始化了。
	getOnce()

	getTwo()
	getTwo()
}

var a int = 0

func getOnce() {
	var one sync.Once

	one.Do(func() {
		a = a + 1
		fmt.Printf("-- %d --\n", a)
	})
}

var one sync.Once
var b int = 2
func getTwo() {
	one.Do(func() {
		b = b + 1
		fmt.Printf("-- %d --\n", b)
	})
}


func GetThree() {
	one.Do(func() {
		a = a + 1
		fmt.Printf("-- %d --\n", a)
	})
}