package sms

import (
	"fmt"
	"net/http"

	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/ports"
	"github.com/pawannn/cashtrack/internal/utils"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilloService struct {
	Env       string
	Client    *twilio.RestClient
	ServiceID string
}

func InitTwilloClient(c *config.CashTrackCfg) ports.SMSRepo {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.SMSAccountSID,
		Password: c.SMSServiceToken,
	})
	return TwilloService{
		Env:       c.ENV,
		Client:    client,
		ServiceID: c.SMSServiceID,
	}
}

func (tR TwilloService) SendOTP(phoneNumber string) utils.CashTrackError {
	if tR.Env != "PROD" {
		return utils.NoErr
	}

	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := tR.Client.VerifyV2.CreateVerification(tR.ServiceID, params)
	if err != nil {
		return utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Unable to send OTP",
			Error:   err,
		}
	}

	if resp != nil && resp.Sid != nil {
		fmt.Printf("OTP sent. SID: %s, Status: %s\n", *resp.Sid, *resp.Status)
	}

	return utils.NoErr
}

func (tR TwilloService) VerifyOTP(phoneNumber string, code string) utils.CashTrackError {
	if tR.Env != "PROD" {
		return utils.NoErr
	}

	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := tR.Client.VerifyV2.CreateVerificationCheck(tR.ServiceID, params)
	if err != nil {
		return utils.CashTrackError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to verify OTP",
			Error:   err,
		}
	}

	if resp != nil && resp.Status != nil {
		if *resp.Status == "approved" {
			return utils.NoErr
		}
		return utils.CashTrackError{
			Code:    http.StatusUnauthorized,
			Message: "Invalid OTP",
			Error:   err,
		}
	}

	return utils.CashTrackError{
		Code:    http.StatusInternalServerError,
		Message: "Invalid response from SMS provider",
		Error:   nil,
	}
}
