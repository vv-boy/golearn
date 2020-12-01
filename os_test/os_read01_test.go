package os

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestRead(t *testing.T) {
	test, _ := os.Open("test.txt")
	fmt.Println("文件的名字为 " + test.Name())

	line := make([]byte, 1000)
	i, _ := test.Read(line)
	fmt.Println("读了个字符" + strconv.Itoa(i))
	fmt.Println(string(line))
	defer test.Close()

	line2 := make([]byte, 1000)
	// 这个 offer 是 line2 的位置了
	i, _ = test.ReadAt(line2, 100)
	fmt.Println(string(line2))

	_, err := test.Write([]byte("test.txt"))

	if err != nil {
		fmt.Println(err)
	}

	test, _ = os.OpenFile("test.txt", os.O_WRONLY | os.O_CREATE, 0)

	_, _ = test.Seek(0, 0)
	_, err= test.Write([]byte("string"))
	if err!= nil {
		fmt.Println(err)
	}

}
