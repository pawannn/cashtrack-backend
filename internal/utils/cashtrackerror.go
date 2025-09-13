package utils

type CashTrackError struct {
	Code    int
	Message string
	Error   error
}

var NoErr = CashTrackError{}
