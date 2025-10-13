package main

import "go.uber.org/zap"

func base() {
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

func sugarLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Info("Info")
	sugar.Infow("Product", "test", 1)
	sugar.Infoln("Info ln")
}

func customLogger() {
	logger, _ := newCustomLogger()
	defer logger.Sync()

	logger.Info("INfo")
}

func main() {
	//base()
	//sugarLogger()
	customLogger()
}
