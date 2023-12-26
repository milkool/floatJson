package jfloat

import (
	"errors"
	"fmt"
)

type JFloat struct {
	Value  float64
	Digits int
}

func NewJFloat(value float64, digits int) JFloat {
	return JFloat{Value: value, Digits: digits}
}

func (j JFloat) MarshalJSON() ([]byte, error) {
	// chechk digits
	if j.Digits < 0 {
		return []byte{}, errors.New("Digits Parameter is less than zero")
	}

	s := fmt.Sprintf("%.*f", j.Digits, j.Value)

	return []byte(s), nil
}
