package test_empty

import (
	"fmt"
	"github.com/novalagung/gubrak/v2"
	"reflect"
	"testing"
)

func TestEmpty(t *testing.T) {
	offers:=[]offer{}
	fmt.Println(gubrak.IsEmpty(offers))

	vvTest := offer{Uuid: "vvtest"}
	fmt.Println(reflect.ValueOf(vvTest).IsZero())
	fmt.Println(10 / 3)
}


type offer struct {
	Uuid string
	OfferName string
}