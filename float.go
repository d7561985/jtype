package jtype

import (
	"strconv"
	"strings"
)

type Float float64

// UnmarshalJSON ...
func (f *Float) UnmarshalJSON(in []byte) error {
	v := strings.Trim(string(in), `"`)
	fl, err := strconv.ParseFloat(v, 0)
	// no prefer errors, just default value
	if err != nil {
		*f = 0
		return nil
	}
	*f = Float(fl)
	return nil
}
