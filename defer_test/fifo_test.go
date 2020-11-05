package defer_test

import (
	"log"
	"strconv"
	"testing"
)

func TestDeferTest(t *testing.T) {
	(func() {
		for i := 0; i < 6; i++ {
			defer log.Println("EDDYCJY" + strconv.Itoa(i) + ".")
		}
		log.Print("end")
	})()

}
