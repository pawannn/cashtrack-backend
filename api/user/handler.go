package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/middlewares"
	cashTrackHttp "github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/pkg/http/payloads"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApi) ValidatePhone(c *gin.Context) {
	reqID := utils.NewUUID()
	var userDetails models.User
	if err := c.BindJSON(&userDetails); err != nil {
		cashTrackHttp.SendResponse(c, reqID, http.StatusInternalServerError, "Unable to read payload", err.Error())
		return
	}

	err := uA.userService.ValidatePhone(userDetails.Phone, userDetails.Country)
	if err != utils.NoErr {
		fmt.Println(err)
		fmt.Println(utils.NoErr)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "OTP Sent to the given phone number", nil)
}

func (uA *UserApi) VerifyPhone(c *gin.Context) {
	reqID := utils.NewUUID()
	var payload payloads.VerifyPhonePayload
	if err := c.BindJSON(&payload); err != nil {
		cashTrackHttp.SendResponse(c, reqID, http.StatusInternalServerError, "Unable to ready payload", nil)
		return
	}

	currency := utils.GetCurrency(payload.Country)
	payload.Currency = currency

	user, err := uA.userService.VerifyPhone(&payload.User, payload.OTP)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}
	token, err := uA.cashtrackEngine.AuthRepo.GenerateUserToken(user.Id)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	response := payloads.VerifyPhoneResponse{
		User:  *user,
		Token: token,
	}
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Phone verified succesfully", response)
}

func (uA *UserApi) GetUser(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, "", err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	user, err := uA.userService.GetUserByID(userID)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Fetched user details successfully", user)
}
