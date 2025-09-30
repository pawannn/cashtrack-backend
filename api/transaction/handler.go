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
		newReqID := utils.NewUUID()
		tA.transactionLogger.Error(newReqID, "Failed to parse context", err.Error)
		cashTrackHttp.SendResponse(c, newReqID, err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	var transaction models.Transaction

	if err := c.BindJSON(&transaction); err != nil {
		tA.transactionLogger.Error(reqID, "Failed to read transaction payload", err)
		cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Unable to read payload", err.Error())
		return
	}

	transaction.UserID = userID
	tA.transactionLogger.Info(reqID, "Recording transaction", transaction)

	txn, err := tA.transactionService.Record(&transaction)
	if err != utils.NoErr {
		tA.transactionLogger.Error(reqID, "Failed to record transaction", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	tA.transactionLogger.Info(reqID, "Transaction recorded successfully", txn.Id)
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Transaction recorded successfully", txn)
}

func (tA *TransactionApi) GetUserTransactions(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		newReqID := utils.NewUUID()
		tA.transactionLogger.Error(newReqID, "Failed to parse context", err.Error)
		cashTrackHttp.SendResponse(c, newReqID, err.Code, err.Message, err.Error)
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
			tA.transactionLogger.Error(reqID, "Invalid 'from' date format", err, fromStr)
			cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Invalid 'from' date format, use YYYY-MM-DD", err.Error())
			return
		}
		from = &parsedFrom
	}

	if toStr != "" {
		parsedTo, err := time.Parse("2006-01-02", toStr)
		if err != nil {
			tA.transactionLogger.Error(reqID, "Invalid 'to' date format", err, toStr)
			cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Invalid 'to' date format, use YYYY-MM-DD", err.Error())
			return
		}
		to = &parsedTo
	}

	tA.transactionLogger.Info(reqID, "Fetching user transactions", "userID", userID, "from", fromStr, "to", toStr)

	transactions, err := tA.transactionService.FilterUserTransactions(userID, from, to)
	if err != utils.NoErr {
		tA.transactionLogger.Error(reqID, "Failed to fetch user transactions", err.Error)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	tA.transactionLogger.Info(reqID, "Fetched user transactions successfully", "count", len(transactions))
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Transactions fetched successfully", transactions)
}

func (tA *TransactionApi) UpdateTransaction(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		newReqID := utils.NewUUID()
		tA.transactionLogger.Error(newReqID, "Failed to parse context", err.Error)
		cashTrackHttp.SendResponse(c, newReqID, err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	var transaction models.Transaction

	if err := c.BindJSON(&transaction); err != nil {
		tA.transactionLogger.Error(reqID, "Failed to read transaction payload for update", err)
		cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Unable to read payload", err.Error())
		return
	}

	transaction.UserID = userID
	tA.transactionLogger.Info(reqID, "Updating transaction", "txnID", transaction.Id, "userID", userID)

	txn, err := tA.transactionService.Update(&transaction)
	if err != utils.NoErr {
		tA.transactionLogger.Error(reqID, "Failed to update transaction", err.Error, "txnID", transaction.Id)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	tA.transactionLogger.Info(reqID, "Transaction updated successfully", "txnID", txn.Id, "userID", userID)
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Transaction updated successfully", txn)
}

func (tA *TransactionApi) DeleteTransaction(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		newReqID := utils.NewUUID()
		tA.transactionLogger.Error(newReqID, "Failed to parse context", err.Error)
		cashTrackHttp.SendResponse(c, newReqID, err.Code, err.Message, err.Error)
		return
	}
	reqID := contextApi.ReqID
	userID := contextApi.UserID

	var transaction models.Transaction

	if err := c.BindJSON(&transaction); err != nil {
		tA.transactionLogger.Error(reqID, "Failed to read transaction payload for delete", err)
		cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Unable to read payload", err.Error())
		return
	}

	transaction.UserID = userID
	tA.transactionLogger.Info(reqID, "Deleting transaction", "txnID", transaction.Id, "userID", userID)

	err = tA.transactionService.Delete(&transaction)
	if err != utils.NoErr {
		tA.transactionLogger.Error(reqID, "Failed to delete transaction", err.Error, "txnID", transaction.Id)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	tA.transactionLogger.Info(reqID, "Transaction deleted successfully", "txnID", transaction.Id, "userID", userID)
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "Transaction deleted successfully", nil)
}

func (tA *TransactionApi) GetUserStats(c *gin.Context) {
	contextApi, err := middlewares.ParseContext(c)
	if err != utils.NoErr {
		newReqID := utils.NewUUID()
		tA.transactionLogger.Error(newReqID, "Failed to parse context", err.Error)
		cashTrackHttp.SendResponse(c, newReqID, err.Code, err.Message, err.Error)
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
			tA.transactionLogger.Error(reqID, "Invalid 'from' date format", err, "input", fromStr)
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
			tA.transactionLogger.Error(reqID, "Invalid 'to' date format", err, "input", toStr)
			cashTrackHttp.SendResponse(c, reqID, http.StatusBadRequest, "Invalid 'to' date format, use YYYY-MM-DD", err.Error())
			return
		}
		to = &parsedTo
	} else {
		to = &endOfMonth
	}

	tA.transactionLogger.Info(reqID, "Fetching user stats", "userID", userID, "from", from.Format("2006-01-02"), "to", to.Format("2006-01-02"))

	stats, err := tA.transactionService.UserStats(userID, from, to)
	if err != utils.NoErr {
		tA.transactionLogger.Error(reqID, "Failed to fetch user stats", err.Error, "userID", userID)
		cashTrackHttp.SendResponse(c, reqID, err.Code, err.Message, err.Error)
		return
	}

	tA.transactionLogger.Info(reqID, "User stats fetched successfully", "userID", userID)
	cashTrackHttp.SendResponse(c, reqID, http.StatusOK, "User statistics fetched successfully", stats)
}
