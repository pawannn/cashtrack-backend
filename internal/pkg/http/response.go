package http

import (
	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/utils"
)

type ApiResponse struct {
	ReqId   string `json:"reqId"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendResponse(c *gin.Context, reqID string, status int, message string, data any) {
	if reqID == "" {
		reqID = utils.NewUUID()
	}
	newApiResponse := ApiResponse{
		ReqId:   reqID,
		Code:    status,
		Message: message,
		Data:    data,
	}
	c.JSON(status, newApiResponse)
}
