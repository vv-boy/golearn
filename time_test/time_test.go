package time_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	local, _ := time.LoadLocation("Local")
	fmt.Println(local.String())
	fmt.Println(time.Local)

	time1 := time.Now()
	fmt.Println(time1)

	t2 := time.Unix(1605196800, 0)
	fmt.Println(t2)
	fmt.Println(t2.Date())

	fmt.Println(t2.Month())

	fmt.Println(int(t2.Month()))
	fmt.Println(t2.Day())

	t3, _ := time.ParseInLocation("2006-01-02 15:04", "2016-12-13 15:34", time.Local)

	fmt.Println(t3.Unix())

	fmt.Println(t3.Format("2006-01-00 00:00:00"))

}



