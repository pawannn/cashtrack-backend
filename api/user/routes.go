package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/domain/services"
	"github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/pkg/logger"
	"github.com/pawannn/cashtrack/internal/ports"
)

type UserApi struct {
	cashtrackEngine *http.CashTrackEngine
	userLogger      *logger.Logger
	userService     *services.UserService
	authService     *ports.AuthRepo
}

func InitUserApi(ctEngine *http.CashTrackEngine, userService *services.UserService, authRepo *ports.AuthRepo) *UserApi {
	_logger := logger.InitNewLogger("user-service")
	return &UserApi{
		cashtrackEngine: ctEngine,
		userLogger:      &_logger,
		userService:     userService,
		authService:     authRepo,
	}
}

func (uA UserApi) AddRoutes() {
	uA.cashtrackEngine.AddV1Routes([]http.CashTrackRoutes{
		{
			Method:      "POST",
			Path:        "/users/validate/phone",
			Handler:     uA.ValidatePhone,
			MiddleWares: []gin.HandlerFunc{},
			Description: "This endpoint validates the phone and sends OTP",
		},
		{
			Method:      "GET",
			Path:        "/users/",
			Handler:     uA.GetUser,
			MiddleWares: []gin.HandlerFunc{},
			Description: "This endpoint gets details of the user",
		},
	})
}
