package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Duration(10 * time.Second))

	time.Sleep(time.Duration(11 * time.Second))
	select {
	case <-timer.C:
		fmt.Println("timer fired early")
	default:
		fmt.Println("default")
	}
}
