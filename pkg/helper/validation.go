package helper

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

func ValidPinfl(pinfl string) error {
	if pinfl == "" {
		return errors.New("error application passport_pinfl requirement body to model")
	}
	pattern := regexp.MustCompile(`^([0-9]{14})$`)

	if !(pattern.MatchString(pinfl)) {
		return errors.New("passport_pinfl must be 14 digits")
	}
	return nil
}

func ValidPassportNumber(number string) error {
	if number == "" {
		return errors.New("error application passport_number requirement body to model")
	}
	pattern := regexp.MustCompile(`^([0-9]{7})$`)

	if !(pattern.MatchString(number)) {
		return errors.New("passport_number must be 7 digits")
	}
	return nil
}

func IsValidUUID(u string) bool {
    _, err := uuid.Parse(u)
    return err == nil
 }
