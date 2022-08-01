package log

import (
	"io"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	LogModeDevelopment = "development"
	LogModeProduction  = "production"
)

type logConfig struct {
	Filename  string
	LogLevel  string
	LogMode   string
	MaxSize   int
	MaxAge    int //days
	Compress  bool
	WithColor bool
	ShowLine  bool
}

func defaultLog() *logConfig {
	return &logConfig{
		Filename:  "",
		LogLevel:  "debug",
		LogMode:   LogModeDevelopment,
		MaxSize:   500,
		MaxAge:    28,   //days
		Compress:  true, // disabled by default
		WithColor: true,
		ShowLine:  true,
	}
}
func (l *logConfig) Enocder() zapcore.Encoder {

	switch l.LogMode {
	case LogModeDevelopment:
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.CallerKey = "caller"
		encoderConfig.MessageKey = "message"
		encoderConfig.LevelKey = "level"
		encoderConfig.TimeKey = "time"
		encoderConfig.NameKey = "logger"
		encoderConfig.CallerKey = "caller"
		if l.WithColor {
			encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}
		return zapcore.NewConsoleEncoder(encoderConfig)
	case LogModeProduction:
		encoderConfig := zap.NewProductionEncoderConfig()
		if l.WithColor {
			encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		}
		return zapcore.NewJSONEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func (l *logConfig) GetLevel() zapcore.Level {
	switch strings.ToLower(l.LogLevel) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func (l *logConfig) GetPriority() zap.LevelEnablerFunc {
	switch strings.ToLower(l.LogLevel) {
	case "debug":
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	case "info":
		return func(level zapcore.Level) bool {
			return level == zap.InfoLevel
		}
	case "warn":
		return func(level zapcore.Level) bool {
			return level == zap.WarnLevel
		}
	case "error":
		return func(level zapcore.Level) bool {
			return level == zap.ErrorLevel
		}
	case "dpanic":
		return func(level zapcore.Level) bool {
			return level == zap.DPanicLevel
		}
	case "panic":
		return func(level zapcore.Level) bool {
			return level == zap.PanicLevel
		}
	case "fatal":
		return func(level zapcore.Level) bool {
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	}
}

func (l *logConfig) Writer() io.Writer {
	switch strings.Trim(l.Filename, " ") {
	case "", "-":
		return os.Stdout
	default:
		return &lumberjack.Logger{
			Filename: l.Filename,
			MaxSize:  l.MaxSize,
			MaxAge:   l.MaxAge,   //days
			Compress: l.Compress, // disabled by default
		}
	}
}

func (l *logConfig) newCore() zapcore.Core {
	logwriter := l.Writer()
	encoder := l.Enocder()

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(logwriter)),
		l.GetLevel(),
	)
	return core
}
