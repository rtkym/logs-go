package optctx

import (
	"context"

	"github.com/rtkym/logs-go"
)

type contextKey uint8

const key contextKey = iota

type OptCtx struct {
	LoggerOptions []rtlog.OptionFunc
}

// NewContext returns a new Context that carries value lc.
func NewContext(ctx context.Context, lc *OptCtx) context.Context {
	return context.WithValue(ctx, key, lc)
}

// NewLogger returns the Context value stored in ctx, if any.
func NewLogger(ctx context.Context) *rtlog.Logger {
	if logContext, ok := ctx.Value(key).(*OptCtx); ok {
		return rtlog.NewWithOption(logContext.LoggerOptions...)
	}

	return rtlog.New()
}
