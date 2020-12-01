package os_test

import (
	"fmt"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	fmt.Println(os.Stat("os_test.go"))

	info, _ := os.Stat("os_test.go")
	fmt.Printf("文件的名字为%s\n", info.Name())
	fmt.Printf("文件的格式为%s\n", info.Mode())

	fmt.Printf("data source为 %s\n", info.Sys())

	file, _ := os.Open("./")
	list1, _ := file.Readdir(0)
	fmt.Println(list1[0])
}
