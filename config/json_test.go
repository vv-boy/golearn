package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int
	FamilyAddress string `json:"family_address"`
}


func TestJsonData(t *testing.T) {
	var c Person
	inputJson, err := ioutil.ReadFile("./test_data/test.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(inputJson, &c)
	if err!= nil {
		fmt.Print(err)
	}

	fmt.Printf("--- t:\n%v\n\n", c)

}