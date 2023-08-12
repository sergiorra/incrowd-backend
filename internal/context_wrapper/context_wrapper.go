package context_wrapper

import (
	"context"

	"github.com/google/uuid"
)

const (
	correlationIdKey = "CorrelationID"
)

func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationIdKey, correlationID)
}

func GetCorrelationID(ctx context.Context) string {
	val := ctx.Value(correlationIdKey)
	if val == nil {
		val = uuid.New().String()
	}

	return val.(string)
}
