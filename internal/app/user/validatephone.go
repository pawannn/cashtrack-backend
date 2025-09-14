package user

import (
	"net/http"
	"strings"

	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApp) ValidatePhone(phone string, country string) utils.CashTrackError {
	countryCode := strings.ToUpper(strings.TrimSpace(country))
	if len(countryCode) != 2 {
		return utils.CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "Please provide two letter country code",
			Error:   nil,
		}
	}
	exist := utils.CheckCountry(countryCode)
	if !exist {
		return utils.CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "Invalid country code",
			Error:   nil,
		}
	}
	parsedPhone, err := utils.FormatPhone(phone, country)
	if err != utils.NoErr {
		return err
	}
	sent, err := uA.cacheRepo.CheckOtpSentNumbers(parsedPhone)
	if err != utils.NoErr {
		return err
	}
	if sent {
		return utils.CashTrackError{
			Code:    http.StatusMethodNotAllowed,
			Message: "Please wait for some time before sending OTP",
			Error:   nil,
		}
	}
	err = uA.smsRepo.SendOTP(parsedPhone)
	if err != utils.NoErr {
		return err
	}
	err = uA.cacheRepo.StoreOtpSentNumbers(parsedPhone)
	return err
}
