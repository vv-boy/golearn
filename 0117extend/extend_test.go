package _117extend

import (
	"fmt"
	"testing"
)

type person struct {
	name string
	age int
}

type man struct {
	person
	gender string
}

type woman struct {
	*person
	gender string
}

func TestMan(t *testing.T) {
	m1 := man{gender:"test"}
	fmt.Println(m1)
	m1.name = "weiwei"
	fmt.Println(m1)

	m2 := &man{gender: "test2"}
	fmt.Println(m2)
	fmt.Println(*m2)
	fmt.Println(m2.gender)
	fmt.Println((*m2).gender)


}



