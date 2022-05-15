package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"
	"mic-trainning-lessons-part4/internal"
	"os"
	"strconv"
)

func GetMQAddr() string {
	return fmt.Sprintf("%s:%d", internal.AppConf.RocketMQConfig.Host,
		internal.AppConf.RocketMQConfig.Port)
}

func ProductMsg(mqAddr, topicName string) {

	newProducer, err := rocketmq.NewProducer(
		producer.WithGroupName("testGroup"),
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{mqAddr})),
		producer.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}
	err = newProducer.Start()
	if err != nil {
		zap.S().Error("生产者错误" + err.Error())
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: topicName,
			Body:  []byte("happy mall" + strconv.Itoa(i)),
		}
		res, err := newProducer.SendSync(context.Background(), msg)
		if err != nil {
			zap.S().Error("发送消息错误" + err.Error())
		} else {
			zap.S().Info("发送消息成功" + res.String() + "==" + res.MsgID)
		}
		err = newProducer.Shutdown()
		if err != nil {
			zap.S().Error("生产着shutdown" + err.Error())
			os.Exit(1)
		}
	}
}

func ConsumMsg(mqAddr, topic string) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{mqAddr})),
	)
	if err != nil {
		panic(err)
	}
	err = c.Subscribe(topic, consumer.MessageSelector{},
		func(ctx context.Context, msgList ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range msgList {
				fmt.Printf("订阅消息%v \n", msgList[i])
			}
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		zap.S().Error("消费消息错误",err.Error())
	}
	err = c.Start()
	if err != nil {
		zap.S().Error(" c.Start",err.Error())
	}
}

func main() {

	mqAddr := GetMQAddr()
	topic := "HappyMall"
	ProductMsg(mqAddr, topic)

}
