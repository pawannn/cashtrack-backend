package ports

type AuthRepo interface {
	GenerateUserToken(userID string) (*string, error)
	ParseUserToken(token string) (*string, error)
}
