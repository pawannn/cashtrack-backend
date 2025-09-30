package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/domain/services"
	"github.com/pawannn/cashtrack/internal/middlewares"
	"github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/pkg/logger"
)

type TransactionApi struct {
	cashtrackEngine    *http.CashTrackEngine
	transactionLogger  *logger.Logger
	transactionService *services.TransactionService
	middlware          middlewares.MiddlewareService
}

func InitTransactionApi(ctEngine *http.CashTrackEngine, transcationService *services.TransactionService) *TransactionApi {
	_logger := logger.InitNewLogger("transcation-service")
	_middleware := middlewares.InitMiddleWares(ctEngine.AuthRepo)

	return &TransactionApi{
		cashtrackEngine:    ctEngine,
		transactionLogger:  &_logger,
		transactionService: transcationService,
		middlware:          _middleware,
	}
}

func (tA TransactionApi) AddRoutes() {
	tA.cashtrackEngine.AddV1Routes([]http.CashTrackRoutes{
		{
			Method:  "GET",
			Path:    "/transactions",
			Handler: tA.GetUserTransactions,
			MiddleWares: []gin.HandlerFunc{
				tA.middlware.AuthUser,
			},
			Description: "This endpoint gets transaction details of a user",
		},
		{
			Method:  "GET",
			Path:    "/transactions/stats",
			Handler: tA.GetUserStats,
			MiddleWares: []gin.HandlerFunc{
				tA.middlware.AuthUser,
			},
			Description: "This endpoint gets transaction stats of a user",
		},
		{
			Method:  "POST",
			Path:    "/transactions",
			Handler: tA.RecordTransaction,
			MiddleWares: []gin.HandlerFunc{
				tA.middlware.AuthUser,
			},
			Description: "This endpoint records transaction details of a user",
		},
		{
			Method:  "PUT",
			Path:    "/transactions",
			Handler: tA.UpdateTransaction,
			MiddleWares: []gin.HandlerFunc{
				tA.middlware.AuthUser,
			},
			Description: "This endpoint updates transaction details of a user",
		},
		{
			Method:  "DELETE",
			Path:    "/transaction",
			Handler: tA.DeleteTransaction,
			MiddleWares: []gin.HandlerFunc{
				tA.middlware.AuthUser,
			},
			Description: "This endpoint deleted transaction of a user",
		},
	})
}
