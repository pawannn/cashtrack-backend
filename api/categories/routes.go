package categories

import (
	"github.com/pawannn/cashtrack/internal/domain/services"
	"github.com/pawannn/cashtrack/internal/middlewares"
	"github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/pkg/logger"
)

type CategoriesApi struct {
	cashtrackEngine   *http.CashTrackEngine
	categoriesLogger  *logger.Logger
	categoriesService *services.CategoriesService
	middlware         middlewares.MiddlewareService
}

func InitCategoriesApi(ctEngine *http.CashTrackEngine, categoriesService *services.CategoriesService) *CategoriesApi {
	_logger := logger.InitNewLogger("category-service")
	_middleware := middlewares.InitMiddleWares(ctEngine.AuthRepo)

	return &CategoriesApi{
		cashtrackEngine:   ctEngine,
		categoriesService: categoriesService,
		categoriesLogger:  &_logger,
		middlware:         _middleware,
	}
}

func (cA *CategoriesApi) AddRoutes() {
	cA.cashtrackEngine.AddV1Routes([]http.CashTrackRoutes{
		{
			Method:      "GET",
			Path:        "/categories",
			Handler:     cA.GetCategories,
			MiddleWares: nil,
			Description: "This Endpoint returns all the categories",
		},
	})
}
