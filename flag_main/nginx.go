package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	ip int
	name string
	languages []string

)

func init() {
	flag.IntVar(&ip, "ip", 1234, "help message for flagname");
	flag.StringVar(&name, "name", "weiwei.yu", "请使用正常的名字")
	flag.Var(newSliceValue([]string{}, &languages), "slice", "以逗号分隔")
}

type sliceValue []string

func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = strings.Split(val, ",")
	return nil
}

func (s *sliceValue) Get() interface{} { return []string(*s) }

func (s *sliceValue) String() string { return strings.Join(*s, ",") }


func main() {
	flag.Parse()
	fmt.Println(ip)
	fmt.Println(name)
	fmt.Println(languages)
}
