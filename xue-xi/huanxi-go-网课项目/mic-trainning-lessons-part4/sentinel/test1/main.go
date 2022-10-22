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

	ch := make(chan struct{})
	for i := 0; i <= 2; i++ {
		go func() {
			for {
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
		}()
	}

	go func() {
		time.Sleep(3 * time.Millisecond)
		_, err = flow.LoadRules([]*flow.Rule{
			{
				ID:                     resName,
				Resource:               resName,
				TokenCalculateStrategy: flow.Direct,
				ControlBehavior:        flow.Reject,
				Threshold:              10,
				StatIntervalInMs:       1000,
			},
		})
		if err != nil {
			panic(err)
		}
	}()
	<-ch
}

/*function main() {

	conf := config.NewDefaultConfig()
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)

	if err != nil {
		panic(err)
	}

	ch := make(chan struct{})

	for i := 0; i <= 2; i++ {
		go function() {
			for {
				entry, blockError := sentinel.Entry(resName, sentinel.WithTrafficType(base.Inbound))
				if blockError != nil {
					fmt.Println("流量太大")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					fmt.Println("限流通过...")
					time.Sleep(time.Duration(rand.Uint64()%10 )* time.Millisecond)
					entry.Exit()
				}
			}
		}()
	}

	go function() {
		time.Sleep(3 * time.Millisecond)
		_, err = flow.LoadRules([]*flow.Rule{
			{
				ID:                     resName,
				Resource:               resName,
				TokenCalculateStrategy: flow.Direct,
				ControlBehavior:        flow.Reject,
				Threshold:              10,
				StatIntervalInMs:       1000,
			},
		})
		if err != nil {
			panic(err)
		}
	}()
	<-ch
}*/

/*function main() {

	err := sentinel.InitDefault()
	if err != nil {
		panic(err)
	}

	var rules []*flow.Rule

	//限流规则
	rule := &flow.Rule{
		ID:                     resName,
		Resource:               "resName",
		TokenCalculateStrategy: flow.Direct,
		ControlBehavior:        flow.Reject,
		Threshold:              10,
		StatIntervalInMs:       1000, // 1秒
	}

	rules = append(rules, rule)
	_, err = flow.LoadRules(rules)
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 20; i++ {
		entry, blockError := sentinel.Entry(resName, sentinel.WithTrafficType(base.Inbound))
		if blockError != nil {
			fmt.Println("流量太大")
			time.Sleep(time.Duration(rand.Uint64() % 5) * time.Millisecond)
		} else {
			fmt.Println("限流通过...")
			time.Sleep(time.Duration(rand.Uint64() % 5) * time.Millisecond)
			entry.Exit()
		}
	}

}*/
