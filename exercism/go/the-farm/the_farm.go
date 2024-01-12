package thefarm

import (
	"errors"
	"strconv"
)

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	fodder, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodder / float64(cows) * factor, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {

	if cows > 0 {
		return DivideFood(fc, cows)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows int
	msg  string
}

func NewInvalidCowsError(cows int) *InvalidCowsError {
	return &InvalidCowsError{cows: cows}
}

func (e *InvalidCowsError) Error() string {
	msg := ""

	if e.cows <= 0 {
		msg = "there are no negative cows"
	}

	if e.cows == 0 {
		msg = "no cows don't need food"
	}
	return strconv.Itoa(e.cows) + " cows are invalid: " + msg
}

func ValidateNumberOfCows(cows int) error {
	if cows <= 0 {
		return NewInvalidCowsError(cows)
	}

	return nil
}
