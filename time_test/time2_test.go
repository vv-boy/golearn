package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime2(t *testing.T) {
	str := time.Now().Add(time.Duration(-3600) * time.Second).Format("2006-01-02 15:04:05")
	fmt.Println(str)
}
