package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"testing"
)

type Conf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func TestReadYaml(t2 *testing.T) {


	var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`
	//var c conf
	//err := yaml.Unmarshal([]byte(data), c)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Print("test")

	var t T

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
}


func TestMyStruct(t *testing.T) {
	c := Conf{}
	yamlFile, err := ioutil.ReadFile("./test_data/test.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println(err.Error())
	}


	fmt.Printf("--- t:\n%v\n\n", c)

}
