package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/domain/services"
	"github.com/pawannn/cashtrack/internal/middlewares"
	"github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/pkg/logger"
)

type UserApi struct {
	cashtrackEngine *http.CashTrackEngine
	userLogger      *logger.Logger
	userService     *services.UserService
	middlware       middlewares.MiddlewareService
}

func InitUserApi(ctEngine *http.CashTrackEngine, userService *services.UserService) *UserApi {
	_logger := logger.InitNewLogger("user-service")
	_middleware := middlewares.InitMiddleWares(ctEngine.AuthRepo)
	return &UserApi{
		cashtrackEngine: ctEngine,
		userLogger:      &_logger,
		userService:     userService,
		middlware:       _middleware,
	}
}

func (uA UserApi) AddRoutes() {
	uA.cashtrackEngine.AddV1Routes([]http.CashTrackRoutes{
		{
			Method:  "GET",
			Path:    "/users",
			Handler: uA.GetUser,
			MiddleWares: []gin.HandlerFunc{
				uA.middlware.AuthUser,
			},
			Description: "This endpoint gets details of the user",
		},
		{
			Method:  "PATCH",
			Path:    "/users",
			Handler: uA.UpdateUser,
			MiddleWares: []gin.HandlerFunc{
				uA.middlware.AuthUser,
			},
			Description: "This endpoint gets details of the user",
		},
		{
			Method:      "POST",
			Path:        "/users/validate/phone",
			Handler:     uA.ValidatePhone,
			MiddleWares: nil,
			Description: "This endpoint validates the phone and sends OTP",
		},
		{
			Method:      "POST",
			Path:        "/users/verify/phone",
			Handler:     uA.VerifyPhone,
			MiddleWares: nil,
			Description: "This enpoint validates the OTP and save user details",
		},
	})
}
