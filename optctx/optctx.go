package optctx

import (
	"context"

	"github.com/rtkym/logs-go"
)

type contextKey uint8

const key contextKey = iota

type OptCtx struct {
	LoggerOptions []logs.OptionFunc
}

// NewContext returns a new Context that carries value lc.
func NewContext(ctx context.Context, lc *OptCtx) context.Context {
	return context.WithValue(ctx, key, lc)
}

// NewLogger returns the Context value stored in ctx, if any.
func NewLogger(ctx context.Context) *logs.Logger {
	if logContext, ok := ctx.Value(key).(*OptCtx); ok {
		return logs.NewWithOption(logContext.LoggerOptions...)
	}

	return logs.New()
}
