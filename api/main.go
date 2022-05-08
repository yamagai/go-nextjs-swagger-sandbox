package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	undo := zap.ReplaceGlobals(logger)

	if err = logger.Sync(); err != nil {
		zap.L().Error(err.Error())
	}
	undo()
}
