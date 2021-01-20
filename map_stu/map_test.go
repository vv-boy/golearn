package map_stu

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	aMap := map[string]int{"one": 1, "two": 2, "three": 3}
	aMap["test"] = 10
	fmt.Println(aMap)
	fmt.Println(aMap)
	var bMap = map[string]int{}
	bMap["test"] = 1

	fmt.Println(bMap)

	for key, _ := range aMap {
		fmt.Println(key)
	}

	for key, _ := range aMap {
		fmt.Println(key)
	}
}
