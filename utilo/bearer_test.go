package utilo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveBearer(t *testing.T) {
	assert.True(t, RemoveBearer("Bearer token") == "token")
	assert.True(t, RemoveBearer(" token") == " token")
}
