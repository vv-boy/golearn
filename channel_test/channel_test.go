package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch1 := make(chan int, 0)
	//ch1 <- 6

	go func() { ch1 <- 10 }()
	go func() { fmt.Println(<-ch1 + 1) }()
	go func() { fmt.Println(<-ch1 + 2) }()
	//ch1 <- 5

	time.Sleep(time.Second * 1)

	//ch1 <- 5

	//time.Sleep(time.Second * 5)
}
