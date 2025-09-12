package ports

type SMSRepo interface {
	SendOTP(phone string) error
	VerifyOTP(phone string, OTP string) (bool, error)
}
