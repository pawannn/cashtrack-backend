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

func InitTransactionApi(ctEngine *http.CashTrackEngine, transactionService *services.TransactionService) *TransactionApi {
	_logger := logger.InitNewLogger("transaction-service")
	_middleware := middlewares.InitMiddleWares(ctEngine.AuthRepo)

	return &TransactionApi{
		cashtrackEngine:    ctEngine,
		transactionLogger:  &_logger,
		transactionService: transactionService,
		middlware:          _middleware,
	}
}

func (tA TransactionApi) AddRoutes() {
	tA.cashtrackEngine.AddV1Routes([]http.CashTrackRoutes{
		{
			Method:  "GET",
			Path:    "/transaction",
			Handler: tA.GetUserTransactions,
			MiddleWares: []gin.HandlerFunc{
				tA.middlware.AuthUser,
			},
			Description: "This endpoint gets transaction details of a user",
		},
		{
			Method:  "GET",
			Path:    "/transaction/stats",
			Handler: tA.GetUserStats,
			MiddleWares: []gin.HandlerFunc{
				tA.middlware.AuthUser,
			},
			Description: "This endpoint gets transaction stats of a user",
		},
		{
			Method:  "POST",
			Path:    "/transaction",
			Handler: tA.RecordTransaction,
			MiddleWares: []gin.HandlerFunc{
				tA.middlware.AuthUser,
			},
			Description: "This endpoint records transaction details of a user",
		},
		{
			Method:  "PUT",
			Path:    "/transaction",
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
			Description: "This endpoint deletes transaction of a user",
		},
	})
}
