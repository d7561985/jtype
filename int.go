package jtype

import (
	"strconv"
	"strings"
)

// A Int represents a JSON number and quoted literal
type Int int

// UnmarshalJSON ...
func (n *Int) UnmarshalJSON(in []byte) error {
	v := strings.Trim(string(in), `"`)
	r, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	*n = Int(r)
	return nil
}
