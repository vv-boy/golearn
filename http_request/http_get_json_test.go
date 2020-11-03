package http_request

import (
	"encoding/json"
	"fmt"
	"github.com/monaco-io/request"
	"testing"
)

type User struct {
	Name string
	Age int
}
func TestHttpGetJson(t *testing.T) {
	client := request.Client{
		URL:    "https://yapi.baidu.com/mock/8826/get/1",
		Method: "GET",
	}
	resp, _ := client.Do()
	var user User
	json.Unmarshal(resp.Data, &user)

	fmt.Printf("-----t %v", user)


}