package logger

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

var (
	Sugar *zap.SugaredLogger
	Log   *zap.Logger

	customTimeFormat string
	CustomFunc       grpc_zap.CodeToLevel
	onceInit         sync.Once
)

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(customTimeFormat))
}

// lvl - global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
func Init(lvl int) error {
	var err error

	onceInit.Do(func() {
		globalLevel := zapcore.Level(lvl)

		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})
		lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= globalLevel && lvl < zapcore.ErrorLevel
		})
		consoleInfos := zapcore.Lock(os.Stdout)
		consoleErrors := zapcore.Lock(os.Stderr)

		ecfg := zap.NewProductionEncoderConfig()
		consoleEncoder := zapcore.NewJSONEncoder(ecfg)

		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleInfos, lowPriority),
		)

		Log = zap.New(core)
		zap.RedirectStdLog(Log)
		Sugar = Log.Sugar()
	})

	return err
}
