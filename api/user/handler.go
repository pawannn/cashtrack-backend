package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/middlewares"
	cashTrackHttp "github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/pkg/http/payloads"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (uA *UserApi) GetUser(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, utils.NewUUID(), err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	uA.userLogger.Info(reqID, "Requested User Info : ", userID)
	user, err := uA.userService.GetUserByID(userID)
	if err != utils.NoErr {
		uA.userLogger.Error(reqID, "Unable to fetch user details", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}
	uA.userLogger.Info(reqID, "Fetched User Info successfully : ", userID)
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Fetched user details successfully", user)
}

func (uA *UserApi) ValidatePhone(c *gin.Context) {
	reqID := utils.NewUUID()
	var userDetails models.User

	if err := c.BindJSON(&userDetails); err != nil {
		uA.userLogger.Error(reqID, "Failed to read payload for phone validation", err)
		cashTrackHttp.SendResponse(c, reqID, http.StatusInternalServerError, "Unable to read payload", err.Error())
		return
	}

	uA.userLogger.Info(reqID, "Validating phone number for user", userDetails.Phone)
	err := uA.userService.ValidatePhone(userDetails.Phone, userDetails.Country)
	if err != utils.NoErr {
		uA.userLogger.Error(reqID, "Phone validation failed", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	uA.userLogger.Info(reqID, "OTP sent successfully for phone number", userDetails.Phone)
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "OTP Sent to the given phone number", nil)
}

func (uA *UserApi) VerifyPhone(c *gin.Context) {
	reqID := utils.NewUUID()
	var payload payloads.VerifyPhonePayload

	if err := c.BindJSON(&payload); err != nil {
		uA.userLogger.Error(reqID, "Failed to read payload for phone verification", err)
		cashTrackHttp.SendResponse(c, reqID, http.StatusInternalServerError, "Unable to read payload", nil)
		return
	}

	uA.userLogger.Info(reqID, "Verifying phone for user", payload.Name)

	err := utils.ValidateUserName(payload.Name)
	if err != utils.NoErr {
		uA.userLogger.Error(reqID, "Username validation failed", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	payload.Currency = utils.GetCurrency(payload.Country)

	user, err := uA.userService.VerifyPhone(&payload.User, payload.OTP)
	if err != utils.NoErr {
		uA.userLogger.Error(reqID, "Phone verification failed", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	uA.userLogger.Info(reqID, "Phone verified successfully", user.Id)

	token, err := uA.cashtrackEngine.AuthRepo.GenerateUserToken(user.Id)
	if err != utils.NoErr {
		uA.userLogger.Error(reqID, "Failed to generate user token", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	response := payloads.VerifyPhoneResponse{
		User:  *user,
		Token: token,
	}

	uA.userLogger.Info(reqID, "User verification and token generation completed", user.Id)
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Phone verified successfully", response)
}

func (uA *UserApi) UpdateUser(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		newReqID := utils.NewUUID()
		uA.userLogger.Error(newReqID, "Failed to parse context", err.Error)
		cashTrackHttp.SendResponse(c, newReqID, err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	var userDetails models.User

	if err := c.BindJSON(&userDetails); err != nil {
		uA.userLogger.Error(reqID, "Failed to read payload for updating user", err)
		cashTrackHttp.SendResponse(c, reqID, http.StatusInternalServerError, "Unable to read payload", err.Error())
		return
	}

	userDetails.Id = userID
	uA.userLogger.Info(reqID, "Updating user details", userID)

	err = utils.ValidateUserName(userDetails.Name)
	if err != utils.NoErr {
		uA.userLogger.Error(reqID, "Username validation failed during update", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	user, err := uA.userService.UpdateUser(&userDetails)
	if err != utils.NoErr {
		uA.userLogger.Error(reqID, "Failed to update user details", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	uA.userLogger.Info(reqID, "User details updated successfully", userID)
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Update user details successfully", user)
}
