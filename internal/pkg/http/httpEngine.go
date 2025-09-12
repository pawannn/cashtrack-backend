package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/pkg/config"
)

type CashTrackRoutes struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	MiddleWares []gin.HandlerFunc
	Description string
}

type CashTrackEngine struct {
	httpEngine *gin.Engine
	Config     *config.CashTrackCfg
}

func InitCashtrackEngine(cfg *config.CashTrackCfg) *CashTrackEngine {
	return &CashTrackEngine{
		httpEngine: gin.Default(),
		Config:     cfg,
	}
}

func (cE *CashTrackEngine) AddRoutes(routes []CashTrackRoutes) {
	for _, route := range routes {
		handlers := append(route.MiddleWares, route.Handler)
		switch route.Method {
		case "GET":
			cE.httpEngine.GET(route.Path, handlers...)
		case "POST":
			cE.httpEngine.POST(route.Path, handlers...)
		case "PUT":
			cE.httpEngine.PUT(route.Path, handlers...)
		case "PATCH":
			cE.httpEngine.PATCH(route.Path, handlers...)
		case "DELETE":
			cE.httpEngine.DELETE(route.Path, handlers...)
		default:
			fmt.Println("Invalid route method : ", route.Method)
			continue
		}
	}
}

func (cE *CashTrackEngine) StartServer() error {
	address := fmt.Sprintf(":%d", cE.Config.Port)
	if err := cE.httpEngine.Run(address); err != nil {
		return err
	}
	return nil
}
