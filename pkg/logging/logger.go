package logging

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
)

type Fields map[string]any

type Logger interface {
	WithContext(ctx context.Context) context.Context
	WithName(name string) Logger
	WithField(key string, value any) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)
	Debugf(msg string, args ...any)
	Infof(msg string, args ...any)
	Warnf(msg string, args ...any)
	Errorf(msg string, args ...any)
	Fatalf(msg string, args ...any)
	Panicf(msg string, args ...any)
	Flush()
}

var (
	initOnce sync.Once
)

const (
	initialSampling    = 100
	thereafterSampling = 100
)

type ZapWrapper struct{ sugaredLogger *zap.SugaredLogger }

var _ Logger = (*ZapWrapper)(nil)

// Debug implements Logger.
func (z *ZapWrapper) Debug(msg string) {
	z.sugaredLogger.Debug(msg)
}

// Debugf implements Logger.
func (z *ZapWrapper) Debugf(msg string, args ...any) {
	z.sugaredLogger.Debugf(msg, args...)
}

// Info implements Logger.
func (z *ZapWrapper) Info(msg string) {
	z.sugaredLogger.Info(msg)
}

// Infof implements Logger.
func (z *ZapWrapper) Infof(msg string, args ...any) {
	z.sugaredLogger.Infof(msg, args...)
}

// Warn implements Logger.
func (z *ZapWrapper) Warn(msg string) {
	z.sugaredLogger.Warn(msg)
}

// Warnf implements Logger.
func (z *ZapWrapper) Warnf(msg string, args ...any) {
	z.sugaredLogger.Warnf(msg, args...)
}

// Error implements Logger.
func (z *ZapWrapper) Error(msg string) {
	z.sugaredLogger.Error(msg)
}

// Errorf implements Logger.
func (z *ZapWrapper) Errorf(msg string, args ...any) {
	z.sugaredLogger.Errorf(msg, args...)
}

// Fatal implements Logger.
func (z *ZapWrapper) Fatal(msg string) {
	z.sugaredLogger.Fatal(msg)
}

// Fatalf implements Logger.
func (z *ZapWrapper) Fatalf(msg string, args ...any) {
	z.sugaredLogger.Fatalf(msg, args...)
}

// Panic implements Logger.
func (z *ZapWrapper) Panic(msg string) {
	z.sugaredLogger.Panic(msg)
}

// Panicf implements Logger.
func (z *ZapWrapper) Panicf(msg string, args ...any) {
	z.sugaredLogger.Panicf(msg, args...)
}

// WithContext implements Logger.
func (z *ZapWrapper) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey, z)
}

// WithName implements Logger.
func (z *ZapWrapper) WithName(name string) Logger {
	sugaredLogger := z.sugaredLogger.Named(name)
	return &ZapWrapper{sugaredLogger: sugaredLogger}
}

// WithError implements Logger.
func (z *ZapWrapper) WithError(err error) Logger {
	sugaredLogger := z.sugaredLogger.With("error", err)
	return &ZapWrapper{sugaredLogger: sugaredLogger}
}

// WithField implements Logger.
func (z *ZapWrapper) WithField(key string, value any) Logger {
	sugaredLogger := z.sugaredLogger.With(key, value)
	return &ZapWrapper{sugaredLogger: sugaredLogger}
}

// WithFields implements Logger.
func (z *ZapWrapper) WithFields(fields Fields) Logger {
	sugaredLogger := z.sugaredLogger
	for key, value := range fields {
		sugaredLogger = sugaredLogger.With(key, value)
	}
	return &ZapWrapper{sugaredLogger: sugaredLogger}
}

// Flush implements Logger.
func (z *ZapWrapper) Flush() {
	if err := z.sugaredLogger.Sync(); err != nil {
		z.sugaredLogger.
			With("error", err).
			Error("failed to sync logger")
	}
}

func InitializationLogger(config *config.Logging, options ...Option) Logger {
	opts := &Options{
		Format:            config.Format,
		Level:             config.Level,
		Name:              config.Name,
		Outputs:           config.Outputs,
		ErrorOutputs:      config.ErrorOutputs,
		DisableCaller:     false,
		DisableStacktrace: false,
		EnableColor:       true,
	}
	for _, opt := range options {
		opts = opt(opts)
	}

	zapLogger, err := newZapLogger(opts)
	if err != nil {
		panic(err)
	}

	return &ZapWrapper{sugaredLogger: zapLogger}
}

func milliSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}

func newZapLogger(opts *Options) (*zap.SugaredLogger, error) {
	var zc *zap.Config
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encodeLevel := zapcore.CapitalLevelEncoder
	if opts.Format == string(ConsoleFormat) && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	initOnce.Do(func() {
		zc = &zap.Config{
			Level:             zap.NewAtomicLevelAt(zapLevel),
			DisableCaller:     opts.DisableCaller,
			DisableStacktrace: opts.DisableStacktrace,
			Sampling: &zap.SamplingConfig{
				Initial:    initialSampling,
				Thereafter: thereafterSampling,
			},
			Encoding: opts.Format,
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "message",
				LevelKey:       "level",
				TimeKey:        "timestamp",
				NameKey:        "logger",
				CallerKey:      "caller",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    encodeLevel,
				EncodeTime:     zapcore.RFC3339TimeEncoder,
				EncodeDuration: milliSecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
				EncodeName:     zapcore.FullNameEncoder,
			},
			OutputPaths:      opts.Outputs,
			ErrorOutputPaths: opts.ErrorOutputs,
		}
	})

	logger, err := zc.Build(zap.AddStacktrace(zap.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	logger = logger.Named(opts.Name)

	return logger.Sugar(), nil
}

func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		logger := ctx.Value(logContextKey)
		if logger != nil {
			return logger.(Logger)
		}
	}

	return nil
}
