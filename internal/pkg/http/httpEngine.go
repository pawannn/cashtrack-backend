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

func (cE *CashTrackEngine) AddV1Routes(routes []CashTrackRoutes) {
	for _, route := range routes {
		handlers := append(route.MiddleWares, route.Handler)
		routePath := "/cashtrack/v0" + route.Path
		switch route.Method {
		case "GET":
			cE.httpEngine.GET(routePath, handlers...)
		case "POST":
			cE.httpEngine.POST(routePath, handlers...)
		case "PUT":
			cE.httpEngine.PUT(routePath, handlers...)
		case "PATCH":
			cE.httpEngine.PATCH(routePath, handlers...)
		case "DELETE":
			cE.httpEngine.DELETE(routePath, handlers...)
		default:
			fmt.Println("Invalid route method : ", route.Method)
			continue
		}

		fmt.Printf("%s : %s : %s\n", route.Method, routePath, route.Description)
	}
}

func (cE *CashTrackEngine) StartServer() error {
	address := fmt.Sprintf(":%d", cE.Config.Port)
	if err := cE.httpEngine.Run(address); err != nil {
		return err
	}
	return nil
}
