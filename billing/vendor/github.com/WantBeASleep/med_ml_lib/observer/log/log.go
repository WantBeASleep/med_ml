package log

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/WantBeASleep/med_ml_lib/slicer"
	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
)

type logKey struct{}
type logValues map[string]any

var key logKey

// вернет значения из контекста
func getContextAttrs(ctx context.Context) logValues {
	ctxLogValues, ok := ctx.Value(key).(logValues)
	if ctxLogValues == nil || !ok {
		return logValues{}
	}
	return ctxLogValues
}

// Добавит в контекст логируемое значение
func WithField(ctx context.Context, k string, v any) context.Context {
	values := getContextAttrs(ctx)
	values[k] = v
	return context.WithValue(ctx, key, values) // несмотря на то что map - ссылочная, надо обновить поле
}

// Добавит в контекст логируемое значение
func WithFields(ctx context.Context, values map[string]any) context.Context {
	ctxvalues := getContextAttrs(ctx)
	for k, v := range values {
		ctxvalues[k] = v
	}
	return context.WithValue(ctx, key, ctxvalues) // несмотря на то что map - ссылочная, надо обновить поле
}

type config struct {
	logger zerolog.Logger
	level  slog.Leveler
}

func defautlConfig() config {
	zerolog.TimeFieldFormat = time.DateTime
	return config{
		logger: zerolog.New(os.Stdout),
		level:  slog.LevelError,
	}
}

// Задаст логгер по умолчанию
//
// опции пресетов должны быть переданы ПЕРВЫМИ
func InitLogger(opts ...LogOption) {
	cfg := defautlConfig()
	for _, o := range opts {
		o.applyOpt(&cfg)
	}

	logger := slog.New(slogzerolog.Option{
		Level:  cfg.level,
		Logger: &cfg.logger,
		AttrFromContext: []func(ctx context.Context) []slog.Attr{
			func(ctx context.Context) []slog.Attr {
				ctxLogValues := getContextAttrs(ctx)
				res := make([]slog.Attr, 0, len(ctxLogValues))
				for k, v := range ctxLogValues {
					res = append(res, slog.Any(k, v))
				}

				return []slog.Attr{slog.Group("context", slicer.PackSlice(res)...)}
			},
		},
	}.NewZerologHandler())
	slog.SetDefault(logger)
}
