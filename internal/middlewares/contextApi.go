package middlewares

import (
	"context"
	"fmt"
)

type ContextType string

var ContextKey ContextType = "CashTrackContext"

type ApiContext struct {
	UserID string `json:"userID"`
}

func AttachContext(ctx context.Context, contextData ApiContext) context.Context {
	c := context.WithValue(ctx, ContextKey, contextData)
	return c
}

func ParseContext(ctx context.Context) (*ApiContext, error) {
	contextData, ok := ctx.Value(ContextKey).(ApiContext)
	if !ok {
		return nil, fmt.Errorf("unable to get data from the given contect")
	}
	return &contextData, nil
}
