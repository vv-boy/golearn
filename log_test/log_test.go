package log_test

import (
	"github.com/apsdehal/go-logger"
	"os"
	"testing"
)

func TestLog1(t *testing.T) {
	log, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}

	// Critically log critical
	log.Critical("This is Critical!")
	log.CriticalF("%+v", err)
	// You can also use fmt compliant naming scheme such as log.Criticalf, log.Panicf etc
	// with small 'f'

	// Debug
	// Since default logging level is Info this won't print anything
	log.Debug("This is Debug!")
	log.DebugF("Here are some numbers: %d %d %f", 10, -3, 3.14)
	// Give the Warning
	log.Warning("This is Warning!")
	log.WarningF("This is Warning!")
	// Show the error
	log.Error("This is Error!")
	log.ErrorF("This is Error!")
	// Notice
	log.Notice("This is Notice!")
	log.NoticeF("%s %s", "This", "is Notice!")
	// Show the info
	log.Info("This is Info!")
	log.InfoF("This is %s!", "Info")

	//log.StackAsError("Message before printing stack");

	// Show warning with format
	log.SetFormat("[%{module}] [%{level}] %{message}")
	log.Warning("This is Warning!") // output: "[test] [WARNING] This is Warning!"
	// Also you can set your format as default format for all new loggers
	logger.SetDefaultFormat("%{message}")
	log.SetLogLevel()
	log2, _ := logger.New("pkg", 1, os.Stdout)
	log2.Error("This is Error!") // output: "This is Error!"

}
