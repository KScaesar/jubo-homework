package util

import (
	"context"

	"github.com/oklog/ulid/v2"
)

func NewUlid() string {
	return ulid.Make().String()
}

type correlationIdKey struct{}

func ContextWithCorrelationId(ctx context.Context, corId string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if corId == "" {
		corId = "empty"
	}
	return context.WithValue(ctx, correlationIdKey{}, corId)
}

func CorrelationIdFromContext(ctx context.Context) string {
	corID, ok := ctx.Value(correlationIdKey{}).(string)
	if ok {
		return corID
	}
	return "no_assign"
}
