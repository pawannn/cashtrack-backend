package utils

import (
	"net/http"
	"regexp"
)

func ValidateUserName(name string) CashTrackError {
	if len(name) == 0 {
		return CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "Please provide a valid username",
			Error:   nil,
		}
	}

	if len(name) >= 15 {
		return CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "Name should be less than 15 characters",
			Error:   nil,
		}
	}

	validName := regexp.MustCompile(`^[A-Za-z ]+$`)
	if !validName.MatchString(name) {
		return CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "name can only contain alphabets and spaces",
			Error:   nil,
		}
	}

	return NoErr
}
