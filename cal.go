package cal

import (
	"errors"
)

func Add(a, b float) (float, error) {
	return a + b, nil
}

func Sub(c, d int) (int, error) {
	return c - d, nil
}

func Mul(a, b int) (int, error) {
	return a * b, nil
}

func Div(a, b int) (int, error) {
	if (b == 0) {
		return -1, errors.New("error: denominator can't be zero")
	}
	return a / b, nil
} //add those