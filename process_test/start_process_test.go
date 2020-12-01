package process_test

import (
	"fmt"
	"os"
	"testing"
)

func TestStart(t *testing.T) {

	output, _ := os.OpenFile("/Users/vv/logs/out.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	procAttr := &os.ProcAttr{
		Files: []*os.File{nil, output, output},
	}
	cwd, err := os.Getwd()
	fmt.Println(cwd)

	if err != nil {
		fmt.Println("look path error:", err)
		os.Exit(1)
	}
	process, err := os.StartProcess("/bin/ls", []string{cwd}, procAttr)

	if err != nil {
		fmt.Println("start process error:", err)
		os.Exit(1)
	}

	procState, err := process.Wait()
	_ = process.Release()
	if err != nil {
		fmt.Println("wait error:", err)
		os.Exit(3)
	}

	fmt.Println(procState.Pid())
	fmt.Println(procState.String())
	cmdData := make([]byte, 21)


	procAttr.Files[1].Seek(0, 0);

	procAttr.Files[1].Read(cmdData)
	fmt.Println(string(cmdData))
	os.Exit(0)
}
