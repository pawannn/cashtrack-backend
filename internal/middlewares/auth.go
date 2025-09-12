package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cashTrackHttp "github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (mS MiddlewareService) AuthUser(c *gin.Context) {
	reqID := utils.NewUUID()
	authroization := c.Request.Header.Get("Authorization")
	if authroization == "" {
		cashTrackHttp.SendResponse(c, reqID, http.StatusNotAcceptable, "Missing Authprization", nil)
		c.Abort()
		return
	}
	userID, err := mS.auth.ParseUserToken(authroization)
	if err != nil {
		cashTrackHttp.SendResponse(c, reqID, http.StatusUnauthorized, "Invalid Authprization", nil)
		c.Abort()
		return
	}

	apiContext := ApiContext{
		ReqID:  reqID,
		UserID: userID,
	}
	AttachContext(c, apiContext)

	c.Next()
}
