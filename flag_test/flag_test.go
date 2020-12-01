package flag_test

import (
	"flag"
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	var ip int
	flag.IntVar(&ip, "ip", 1234, "help message for flagname");
	fmt.Println(ip)

}


