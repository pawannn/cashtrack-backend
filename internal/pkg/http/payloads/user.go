package payloads

import "github.com/pawannn/cashtrack/internal/domain/models"

type VerifyPhonePayload struct {
	models.User
	OTP string `json:"otp"`
}

type VerifyPhoneResponse struct {
	models.User
	Token string `json:"token"`
}
