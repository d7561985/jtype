package jtype

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGames(t *testing.T) {
	type A struct {
		A Int `json:"a"`
		B Int `json:"b"`
	}
	a := &A{}
	var q = []byte(`{"a":1, "b":"2"}`)
	err := json.Unmarshal(q, a)
	assert.NoError(t, err)
	assert.Equal(t, Int(1), a.A)
	assert.Equal(t, Int(2), a.B)
}
