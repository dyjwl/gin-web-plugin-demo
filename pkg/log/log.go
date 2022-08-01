package log

import (
	"go.uber.org/zap"
)

var (
	_loger *zap.Logger
)

type Option func(*logConfig)

func WithFilename(filename string) Option {
	return func(l *logConfig) {
		l.Filename = filename
	}
}

func WithLogLevel(logLevel string) Option {
	return func(l *logConfig) {
		l.LogLevel = logLevel
	}
}

func WithLogMode(logMode string) Option {
	return func(l *logConfig) {
		l.LogMode = logMode
	}
}

func WithMaxSize(maxSize int) Option {
	return func(l *logConfig) {
		l.MaxSize = maxSize
	}
}

func WithMaxAge(maxAge int) Option {
	return func(l *logConfig) {
		l.MaxAge = maxAge
	}
}

func WithCompress(compress bool) Option {
	return func(l *logConfig) {
		l.Compress = compress
	}
}

func WithColor(withColor bool) Option {
	return func(l *logConfig) {
		l.WithColor = withColor
	}
}

func WithShowLine(showLine bool) Option {
	return func(l *logConfig) {
		l.ShowLine = showLine
	}
}

func InitLog(opts ...Option) error {
	dlog := defaultLog()
	for _, opt := range opts {
		opt(dlog)
	}
	core := dlog.newCore()
	_loger = zap.New(core)
	if dlog.ShowLine {
		_loger.WithOptions(zap.AddCaller())
	}
	return nil
}

// Debug output log
func Debug(msg string, fields ...zap.Field) {
	_loger.Debug(msg, fields...)
}

// Info output log
func Info(msg string, fields ...zap.Field) {
	_loger.Info(msg, fields...)
}

// Warn output log
func Warn(msg string, fields ...zap.Field) {
	_loger.Warn(msg, fields...)
}

// Error output log
func Error(msg string, fields ...zap.Field) {
	_loger.Error(msg, fields...)
}

// Panic output panic
func Panic(msg string, fields ...zap.Field) {
	_loger.Panic(msg, fields...)
}

// Fatal output log
func Fatal(msg string, fields ...zap.Field) {
	_loger.Fatal(msg, fields...)
}
