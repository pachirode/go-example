package main

import "go.uber.org/zap"

func main() {
	{
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		logger.Info("Prod Test Info", zap.Int("int", 3))
	}

	{
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()

		logger.Info("Dev Test Info", zap.Int("int", 3))
	}
}
