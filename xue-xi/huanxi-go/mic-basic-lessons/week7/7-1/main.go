package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	production, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer production.Sync()
	sugar := production.Sugar()

	e := "出错了"
	sugar.Infof("err:%s", e)
	sugar.Info("又出错了", "err", e, "time:", time.Now().Unix())

	production.Info("logger info test", zap.String("name err", e), zap.Int("price err", 12))

}
