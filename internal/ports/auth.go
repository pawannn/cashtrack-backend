package ports

import "github.com/pawannn/cashtrack/internal/utils"

type AuthRepo interface {
	GenerateUserToken(userID string) (string, utils.CashTrackError)
	ParseUserToken(token string) (string, utils.CashTrackError)
}
