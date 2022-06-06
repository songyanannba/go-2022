package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"math/rand"
	"time"
)

const resName = "ww"

func main() {
	//预热
	err := sentinel.InitDefault()
	if err != nil {
		panic(err)
	}

	var (
		all     int
		through int
		block   int
		ch      = make(chan struct{})
	)

	var rules []*flow.Rule
	//限流规则
	rule := &flow.Rule{
		ID:                     resName,
		Resource:               resName,
		TokenCalculateStrategy: flow.WarmUp,
		ControlBehavior:        flow.Reject,
		Threshold:              1000,
		WarmUpPeriodSec:        30,
	}

	rules = append(rules, rule)
	_, err = flow.LoadRules(rules)
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 10; i++ {
		go func() {
			for {
				entry, blockError := sentinel.Entry(resName, sentinel.WithTrafficType(base.Inbound))
				all++
				if blockError != nil {
					block++
					//fmt.Println("流量太大")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					through++
					//fmt.Println("限流通过...")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					entry.Exit()
				}
			}
		}()
	}

	go func() {
		var oldAll, oldThrough, oldBlock int
		fmt.Println("qwer.....")
		for {
			a := all - oldAll
			oldAll = all

			t := through - oldThrough
			oldThrough = through

			b := block - oldBlock
			oldBlock = block

			time.Sleep(time.Second * 1)
			fmt.Printf("all :%d , thougth:%d ,block:%d \n", a, t, b)
		}
	}()

	<-ch
}
