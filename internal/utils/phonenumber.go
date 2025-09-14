package utils

import (
	"net/http"

	"github.com/nyaruka/phonenumbers"
)

func FormatPhone(phone string, country string) (string, CashTrackError) {
	parsedPhone, err := phonenumbers.Parse(phone, country)
	if err != nil {
		return "", CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "Invalid Phone Number",
			Error:   err,
		}
	}

	formattedPhone := phonenumbers.Format(parsedPhone, phonenumbers.E164)
	return formattedPhone, NoErr
}
