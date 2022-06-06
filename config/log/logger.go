package log

import (
	defaultLog "log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetupLogger configure a global zap logger.
// TODO: Include env vars.
func SetupLogger() {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	log, err := config.Build()
	if err != nil {
		defaultLog.Fatalf("Couldn't initialize logger: %v", err)
	}

	_ = zap.ReplaceGlobals(log)
}
