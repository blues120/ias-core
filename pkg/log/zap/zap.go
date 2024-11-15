package zap

import (
	"fmt"
	"os"
	"strings"

	"gitlab.ctyuncdn.cn/ias/ias-core/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ log.Logger = (*ZapLogger)(nil)

func NewLogger(c *conf.Log) log.Logger {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	level, err := zap.ParseAtomicLevel(c.Level)
	if err != nil {
		panic("logger level parse error: " + err.Error())
	}
	return NewZapLogger(
		c,
		encoder,
		level,
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.Development(),
	)
}

// ZapLogger is a logger impl.
type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

// NewZapLogger return a zap logger.
func NewZapLogger(c *conf.Log, encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	var core zapcore.Core

	// 根据配置文件选择日志打印方式
	switch c.Mode {
	case "console":
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台
			level,
		)
	case "file":
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getFileLogWriter(c.Rotation))), // 打印到控制台和文件
			level,
		)
		if c.SeparateErrorLog {
			// 错误日志额外生成一个文件
			rotation := *c.Rotation
			if strings.HasSuffix(rotation.Filename, ".log") {
				rotation.Filename = strings.Replace(rotation.Filename, ".log", ".err.log", 1)
			} else {
				rotation.Filename += ".err"
			}
			errCore := zapcore.NewCore(
				zapcore.NewJSONEncoder(encoder),
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(getFileLogWriter(&rotation))),
				zap.ErrorLevel,
			)
			core = zapcore.NewTee(core, errCore)
		}
	default:
		panic("unsupported log mode")
	}

	zapLogger := zap.New(core, opts...)
	return &ZapLogger{log: zapLogger, Sync: zapLogger.Sync}
}

// Log Implementation of logger interface.
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	case log.LevelFatal:
		l.log.Fatal("", data...)
	}
	return nil
}

// getFileLogWriter 获取文件日志记录器
func getFileLogWriter(c *conf.Log_Rotation) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   c.Filename,        // 指定日志存储位置
		MaxSize:    int(c.MaxSizeMb),  // 日志的最大大小（M）
		MaxBackups: int(c.MaxBackups), // 日志的最大保存数量
		MaxAge:     int(c.MaxAge),     // 日志文件存储最大天数
		Compress:   c.Compress,        // 是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
