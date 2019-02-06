package jtype

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloat_UnmarshalJSON(t *testing.T) {
	V := []struct {
		IN  string
		Res Float
	}{
		{"0", Float(0)},
		{"0.0000000", Float(0)},
		{"1", Float(1)},
		{"1.1", Float(1.1)},
		{`"0"`, Float(0)},
		{"\"0.0000000\"", Float(0)},
		{"\"1\"", Float(1)},
		{"\"1.1\"", Float(1.1)},
		{`""`, Float(0)},
		{`"qqq"`, Float(0)},
	}

	for _, c := range V {
		t.Run(c.IN, func(t *testing.T) {
			var out Float
			err := json.Unmarshal([]byte(c.IN), &out)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, out, c.Res)
		})
	}
}
