package http_request

import (
	"fmt"
	"github.com/monaco-io/request"
	"log"
	"reflect"
	"stu/http_request/data_json"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestHttpGetJsonArray(t *testing.T) {
	client := request.Client{
		URL:    "https://yapi.baidu.com/mock/8826/get/1",
		Method: "GET",
	}
	resp, err := client.Do()
	fmt.Println(reflect.TypeOf(resp.Data))
	log.Println(resp.Code, string(resp.Data), err)
	//var user data_json.UserList
	////lexer := jlexer.Lexer{Data: resp.Data}
	//user.UnmarshalJSON(resp.Data)

	//if lexer.Error() != nil {
	//	panic(lexer.Error())
	//}

	var user data_json.User
	//json.Unmarshal(resp.Data, &user)
	err = user.UnmarshalJSON(resp.Data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("-----t %v", user)

}
