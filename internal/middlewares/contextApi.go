package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var ContextKey string = "CashTrackContext"

type ApiContext struct {
	ReqID  string `json:"req_id"`
	UserID string `json:"userID"`
}

func AttachContext(ctx *gin.Context, contextData ApiContext) {
	ctx.Set(ContextKey, contextData)
}

func ParseContext(ctx *gin.Context) (*ApiContext, error) {
	contextData, ok := ctx.Value(ContextKey).(ApiContext)
	if !ok {
		return nil, fmt.Errorf("unable to get data from the given contect")
	}
	return &contextData, nil
}
