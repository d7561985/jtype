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
	if err != nil {
		return err
	}
	*f = Float(fl)
	return nil
}
