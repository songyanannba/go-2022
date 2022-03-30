package main

import (
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"time"
)

func main() {
	cft := jaegerCfg.Configuration{
		Sampler:             &jaegerCfg.SamplerConfig{
			Type:                     jaeger.SamplerTypeConst,
			Param:                    1,
		},
		Reporter:            &jaegerCfg.ReporterConfig{
			LogSpans:                   true,
			LocalAgentHostPort:         "192.168.1.5:6831",
		},
		ServiceName: "happyMall",
	}

	tracer, closer, err := cft.NewTracer(jaegerCfg.Logger(jaeger.StdLogger))
	defer closer.Close()
	if err != nil {
		panic(err)
	}
	span := tracer.StartSpan("account_web")
	span.Finish()
	time.Sleep(1 * time.Second)
}
