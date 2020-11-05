package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"stu/package/channel/tutorial2"
	"testing"
)

func TestPackage2(t *testing.T) {
	res :=tutorial2.Add(1, 3)
	fmt.Printf("------ %d\n", res)
	assert.EqualValues(t, res, 4)
}

