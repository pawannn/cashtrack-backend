package ports

import "github.com/pawannn/cashtrack/internal/utils"

type SMSRepo interface {
	SendOTP(phone string) utils.CashTrackError
	VerifyOTP(phone string, OTP string) utils.CashTrackError
}
