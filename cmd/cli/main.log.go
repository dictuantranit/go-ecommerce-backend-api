package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	/* sugar := zap.NewExample().Sugar()
	sugar.Infof("Hello name:%s, age:%d", "TipGo", 40)

	// logger
	logger := zap.NewExample()
	logger.Info("Hello", zap.String("name", "TipGo"), zap.Int("age", 40)) */

	logger := zap.NewExample()
	logger.Info("Hello NewExample")

	// Development
	logger, _ = zap.NewDevelopment()
	logger.Info("Hello NewDevelopment")

	// Production
	logger, _ = zap.NewProduction()
	logger.Info("Hello NewProduction")

}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
