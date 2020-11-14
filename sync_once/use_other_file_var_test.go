package sync_once

import (
	"fmt"
	"testing"
)

func TestOther(t *testing.T) {
	getOne()
	getOne()

	GetThree()
	GetThree()
}

func getOne() {
	one.Do(func() {
		a = a+1
		fmt.Printf("--- %d ---", a);
	});
}


