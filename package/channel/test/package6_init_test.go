package test

import (
	"github.com/stretchr/testify/assert"
	test "stu/package/channel/tutorial"
	"testing"
)


func TestPackage6(t *testing.T) {
	res :=test.Add(1, 2)
	assert.Equal(t, res, 3)
}
