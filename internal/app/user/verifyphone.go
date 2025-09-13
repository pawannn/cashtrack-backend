package user

import "github.com/pawannn/cashtrack/internal/utils"

func (uA *UserApp) VerifyPhone(phone string, country string, OTP string) (bool, utils.CashTrackError) {
	return false, utils.NoErr
}
