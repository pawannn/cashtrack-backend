package user

import (
	"net/http"
	"strings"

	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApp) VerifyPhone(phone string, country string, OTP string) (bool, utils.CashTrackError) {
	countryCode := strings.ToUpper(strings.TrimSpace(country))
	if len(countryCode) != 2 {
		return false, utils.CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "Please provide two letter country code",
			Error:   nil,
		}
	}

	exist := utils.CheckCountry(countryCode)
	if !exist {
		return false, utils.CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "Invalid country code",
			Error:   nil,
		}
	}

	parsedPhone, err := utils.FormatPhone(phone, country)
	if err != utils.NoErr {
		return false, err
	}

	sent, err := uA.cacheRepo.CheckOtpSentNumbers(parsedPhone)
	if err != utils.NoErr {
		return false, err
	}
	if !sent {
		return false, utils.CashTrackError{
			Code:    404,
			Message: "The OTP for the given number has expired or not sent",
			Error:   nil,
		}
	}
	err = uA.smsRepo.VerifyOTP(parsedPhone, OTP)
	if err != utils.NoErr {
		return false, err
	}
	return true, utils.NoErr
}
