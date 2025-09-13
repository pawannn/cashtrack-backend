package user

import "github.com/pawannn/cashtrack/internal/utils"

func (uA *UserApp) ValidatePhone(phone string, country string) utils.CashTrackError {
	return utils.NoErr
}
