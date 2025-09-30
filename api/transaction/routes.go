package transaction

import (
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
