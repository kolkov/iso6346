package iso6346

import (
	"errors"
	"strconv"
	"strings"
)

var ErrValueNotIso6346 = errors.New("iso6346: must be valid number by iso6346 algorithm")
var ErrValueNotLen11 = errors.New("iso6346: must be 11 char length")

// Validate check number is valid or not based on iso6346 algorithm
func Validate(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return ErrValueNotIso6346
	}
	if len(s) != 11 {
		return ErrValueNotLen11
	}
	part := s[0:10]
	checkDigitCalc, err := CalcCheckDigit(part)
	if err != nil {
		return err
	}
	checkDigit, err := strconv.Atoi(s[10:11])
	if err != nil {
		return err
	}
	if checkDigitCalc == 10 {
		checkDigitCalc = 0
	}
	result := checkDigitCalc == checkDigit
	if !result {
		return ErrValueNotIso6346
	}
	return nil
}

var ErrValueNotLen10 = errors.New("iso6346: must be 10 char")

// CalcCheckDigit calculates check digit for iso6346 algorithm.
func CalcCheckDigit(code string) (int, error) {
	if len(code) != 10 {
		return 0, ErrValueNotLen10
	}
	n := 0.0
	d := 0.5
	for _, character := range code {
		d *= 2
		n += d * float64(strings.IndexRune("0123456789A?BCDEFGHIJK?LMNOPQRSTU?VWXYZ", character))
	}
	return int(n) - int(n/11)*11, nil
}

// Generate return the number with check digit
func Generate(source string) (string, error) {
	cd, err := CalcCheckDigit(source)
	return source + strconv.Itoa(cd), err
}
