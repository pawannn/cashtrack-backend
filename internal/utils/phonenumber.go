package utils

import (
	"net/http"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

func FormatPhone(phone string, country string) (string, CashTrackError) {
	countryCode := strings.ToUpper(strings.TrimSpace(country))
	if _, ok := countryCurrency[countryCode]; !ok {
		return "", CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "",
			Error:   nil,
		}
	}
	parsedPhone, err := phonenumbers.Parse(phone, countryCode)
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
