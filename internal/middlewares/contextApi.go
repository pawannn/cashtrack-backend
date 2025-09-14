package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/utils"
)

var ContextKey string = "CashTrackContext"

type ApiContext struct {
	ReqID  string `json:"req_id"`
	UserID string `json:"userID"`
}

func AttachContext(ctx *gin.Context, contextData ApiContext) {
	ctx.Set(ContextKey, contextData)
}

func ParseContext(ctx *gin.Context) (*ApiContext, utils.CashTrackError) {
	contextData, ok := ctx.Value(ContextKey).(ApiContext)
	if !ok {
		return nil, utils.CashTrackError{
			Code:    http.StatusBadRequest,
			Message: "unable to get data from the given contect",
			Error:   nil,
		}
	}
	return &contextData, utils.NoErr
}
