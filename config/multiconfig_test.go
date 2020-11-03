package config

import (
	"fmt"
	"github.com/koding/multiconfig"
	"testing"
)

type Woman struct {
	Name string
	Age  int
	FamilyAddress string
}


func TestMultiConfig(t *testing.T) {
	m := multiconfig.NewWithPath("./test_data/test.json")
	m = multiconfig.New()

	var c Woman
	err := m.Load(&c)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%s,%d", c.Name, c.Age)

}