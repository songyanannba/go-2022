package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
	"math/rand"
	"time"
)

const resName = "cart-order"

func main() {
	conf := config.NewDefaultConfig()
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		panic(err)
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			ID:                     resName,
			Resource:               resName,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Throttling,
			Threshold:              1,
			StatIntervalInMs:       1000,
		},
	})
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 20; i++ {
		entry, blockError := sentinel.Entry(resName, sentinel.WithTrafficType(base.Inbound))
		if blockError != nil {
			fmt.Println("流量太大")
			time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
		} else {
			fmt.Println("限流通过...")
			time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
			entry.Exit()
		}


	}

}
