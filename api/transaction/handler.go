package transaction

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pawannn/cashtrack/internal/domain/models"
	"github.com/pawannn/cashtrack/internal/middlewares"
	cashTrackHttp "github.com/pawannn/cashtrack/internal/pkg/http"
	"github.com/pawannn/cashtrack/internal/utils"
)

func (tA *TransactionApi) RecordTransaction(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, utils.NewUUID(), err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Unable to read payload", err.Error())
		return
	}

	transaction.UserID = userID

	txn, err := tA.transactionService.Record(&transaction)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Transaction recorded successfully", txn)
}

func (tA *TransactionApi) GetUserTransactions(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, utils.NewUUID(), err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	fromStr := c.Query("from")
	toStr := c.Query("to")
	var from, to *time.Time

	if fromStr != "" {
		parsedFrom, err := time.Parse("2006-01-02", fromStr)
		if err != nil {
			cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Invalid 'from' date format, use YYYY-MM-DD", err.Error())
			return
		}
		from = &parsedFrom
	}
	if toStr != "" {
		parsedTo, err := time.Parse("2006-01-02", toStr)
		if err != nil {
			cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Invalid 'to' date format, use YYYY-MM-DD", err.Error())
			return
		}
		to = &parsedTo
	}

	transactions, err := tA.transactionService.FilterUserTransactions(userID, from, to)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Transactions fetched successfully", transactions)
}

func (tA *TransactionApi) UpdateTransaction(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, utils.NewUUID(), err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Unable to read payload", err.Error())
		return
	}

	transaction.UserID = userID

	txn, err := tA.transactionService.Update(&transaction)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Transaction updated successfully", txn)
}

func (tA *TransactionApi) DeleteTransaction(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, utils.NewUUID(), err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	var transaction models.Transaction
	if err := c.BindJSON(&transaction); err != nil {
		cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Unable to read payload", err.Error())
		return
	}

	transaction.UserID = userID

	err = tA.transactionService.Delete(&transaction)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Transaction deleted successfully", nil)
}

func (tA *TransactionApi) GetUserStats(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, utils.NewUUID(), err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	fromStr := c.Query("from")
	toStr := c.Query("to")
	var from, to *time.Time

	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1).Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	if fromStr != "" {
		parsedFrom, err := time.Parse("2006-01-02", fromStr)
		if err != nil {
			cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Invalid 'from' date format, use YYYY-MM-DD", err.Error())
			return
		}
		from = &parsedFrom
	} else {
		from = &startOfMonth
	}

	if toStr != "" {
		parsedTo, err := time.Parse("2006-01-02", toStr)
		if err != nil {
			cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Invalid 'to' date format, use YYYY-MM-DD", err.Error())
			return
		}
		to = &parsedTo
	} else {
		to = &endOfMonth
	}

	stats, err := tA.transactionService.UserStats(userID, from, to)
	if err != utils.NoErr {
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "User statistics fetched successfully", stats)
}
