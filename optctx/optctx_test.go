package optctx_test

import (
	"bytes"
	"context"
	"testing"

	rtlog "github.com/rtkym/logs-go"
	"github.com/rtkym/logs-go/optctx"
	"github.com/stretchr/testify/assert"
)

func TestFromContext(t *testing.T) {
	t.Run("LogContextなし", func(t *testing.T) {
		ctx := context.Background()

		logger := optctx.NewLogger(ctx)
		logger.Info("test")

		assert.NotNil(t, logger)
	})

	t.Run("LogContextあり、Optionsなし", func(t *testing.T) {
		ctx := context.Background()
		ctx = optctx.NewContext(ctx, &optctx.OptCtx{})

		logger := optctx.NewLogger(ctx)
		logger.Info("test")

		assert.NotNil(t, logger)
	})

	t.Run("LogContextあり、Optionsあり", func(t *testing.T) {
		ctx := context.Background()
		buf := &bytes.Buffer{}
		ctx = optctx.NewContext(ctx, &optctx.OptCtx{LoggerOptions: []rtlog.OptionFunc{func(opt *rtlog.Option) { opt.Writer = buf }}})

		logger := optctx.NewLogger(ctx)
		logger.Info("test")

		assert.NotNil(t, logger)
		assert.Contains(t, buf.String(), `"message":"test"`)
	})
}
