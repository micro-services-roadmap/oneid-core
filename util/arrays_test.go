package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {

	assert.False(t, Contains[int]([]int{}, 1))
	assert.True(t, Contains[int]([]int{1, 2}, 1))
	assert.False(t, Contains[int]([]int{1, 2}, 3))
	assert.True(t, Contains[string]([]string{"1", "2"}, "1"))

}
