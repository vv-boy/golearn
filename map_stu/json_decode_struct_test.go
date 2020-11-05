package map_stu

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestMap1(t *testing.T) {
	m := map[string]interface{}{}
	user := User{}
	user.Name = "vi_test"
	user.Age = 12
	m["weiwei"] = user
	str, _ := json.Marshal(m)
	fmt.Println(string(str))
}
