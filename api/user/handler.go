package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/domain/models"
	cashTrackHttp "github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApi) ValidatePhone(c *gin.Context) {
	var userDetails models.User
	if err := c.BindJSON(&userDetails); err != nil {
		cashTrackHttp.SendResponse(c, utils.NewUUID(), http.StatusInternalServerError, "Unable to read payload", nil)
		return
	}
	err := uA.userService.ValidatePhone(userDetails.Phone, userDetails.Country)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, utils.NewUUID(), err.Code, err.Message, err.Error)
		return
	}
}

func (uA *UserApi) VerifyPhone(c *gin.Context) {

}

func (uA *UserApi) GetUser(c *gin.Context) {

}
