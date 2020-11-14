package sort_stu

import (
	"fmt"
	"sort"
	"testing"
)

type User struct {
	Name string
	Age  int
}

type Users []User

func (s Users) Len() int {
	return len(s)
}

func (s Users) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}


func (s Users) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestUser(t *testing.T) {
	users := Users{User{"weiwei", 21},User{"Yellow", 21},User{"vv_test", 25}}
	sort.Sort(users)

	fmt.Println(users)

}
