package shared

import "go.uber.org/zap"

func Log(log string) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("oh no......")
	}
	logger.Info(log)
	defer logger.Sync()
}
