package sms

import (
	"fmt"

	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/ports"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilloService struct {
	Client    *twilio.RestClient
	ServiceID string
}

func InitTwilloClient(c config.CashTrackCfg) ports.SMSRepo {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.SMSAccountSID,
		Password: c.SMSServiceToken,
	})
	return TwilloService{
		Client:    client,
		ServiceID: c.SMSServiceID,
	}
}

func (tR TwilloService) SendOTP(phoneNumber string) error {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := tR.Client.VerifyV2.CreateVerification(tR.ServiceID, params)
	if err != nil {
		return fmt.Errorf("failed to send OTP: %w", err)
	}

	if resp != nil && resp.Sid != nil {
		fmt.Printf("OTP sent. SID: %s, Status: %s\n", *resp.Sid, *resp.Status)
	}

	return nil
}

func (tR TwilloService) VerifyOTP(phoneNumber string, code string) (bool, error) {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := tR.Client.VerifyV2.CreateVerificationCheck(tR.ServiceID, params)
	if err != nil {
		return false, fmt.Errorf("failed to verify OTP: %w", err)
	}

	if resp != nil && resp.Status != nil {
		if *resp.Status == "approved" {
			return true, nil
		}
		return false, nil
	}

	return false, fmt.Errorf("invalid response from Twilio")
}
