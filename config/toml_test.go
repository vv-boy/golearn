package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"testing"
)


type Student struct {
	Name string `toml:"name"`
	Age  int
	FamilyAddress string `toml:"family_address"`
}

type ConfToml struct {
	Student Student
}


func TestToml(t *testing.T) {
	tomlFile, err := ioutil.ReadFile("./test_data/test.toml")
	if err != nil {
		fmt.Println(err.Error())
	}
	c := ConfToml{}
	err = toml.Unmarshal(tomlFile, &c)
	if err != nil {
		fmt.Printf("--- error:\n%v\n\n", err)
	}

	fmt.Printf("--- t:\n%v\n\n", c.Student)
}
