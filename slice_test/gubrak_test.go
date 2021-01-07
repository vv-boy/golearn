package slice

import (
	"fmt"
	"github.com/novalagung/gubrak/v2"
	"reflect"
	"testing"
)

func TestGubrak(t *testing.T) {
	names := []string{"vvtest", "vvtest2", "vvtest3"}
	result := gubrak.From(names).Result().([]string)
	fmt.Println(reflect.TypeOf(result))
	fmt.Println(result[100])
}