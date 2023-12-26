package jfloat

import (
	"errors"
	"fmt"
)

type JFloat struct {
	Value  float64
	Digits int
}

func (j JFloat) MarshalJSON() ([]byte, error) {
	// chechk digits
	if j.Digits < 0 {
		return []byte{}, errors.New("Digits Parameter is less than zero")
	}

	s := fmt.Sprintf("%.*f", j.Digits, j.Value)

	return []byte(s), nil
}
