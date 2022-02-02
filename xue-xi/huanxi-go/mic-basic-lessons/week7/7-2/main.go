package main

import (
	"go.uber.org/zap"
	"time"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = append(config.OutputPaths, "./week7/7-2/7-2.log")
	return config.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()
	defer sugar.Sync()

	e := "出错了..."
	sugar.Infof("err:%s\n", e)
	sugar.Info("又出错了", "err", e, "time:", time.Now().Unix())
}
