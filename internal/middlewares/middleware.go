package middlewares

import "github.com/pawannn/cashtrack/internal/ports"

type MiddlewareService struct {
	auth ports.AuthRepo
}

func InitMiddleWares(authRepo ports.AuthRepo) MiddlewareService {
	return MiddlewareService{
		auth: authRepo,
	}
}
