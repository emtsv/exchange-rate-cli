package cbr

import (
	"errors"
	"strconv"
	"strings"
)

var ErrNomin = errors.New("номинал равен нулю")
var ErrVal = errors.New("валюта")
var ErrEmpty = errors.New("пустое поле")

func (v Valute) RateRUB() (float64, error) {
	if v.Nominal == 0 {
		return 0, ErrEmpty
	}

	clean := strings.ReplaceAll(v.Value, ",", ".")
	valueFloat, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return 0, ErrVal
	}

	return valueFloat / float64(v.Nominal), nil
}
