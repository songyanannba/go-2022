package network

import (
	"encoding/binary"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"github.com/phuhao00/greatestworks-proto/gen/messageId"
	"net"
	"time"
	"github.com/phuhao00/network"
)

type ClientPacket struct {

}

type Client struct {
	cli             *network.Client
	inputHandlers   map[string]InputHandler
	messageHandlers map[messageId.MessageId]MessageHandler
	console         *ClientConsole
	chInput         chan *InputParam
}

func NewClient() *Client {
	c := &Client{
		cli:             network.NewClient(":8023", 200, logger.Logger),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[messageId.MessageId]MessageHandler{},
		console:         NewClientConsole(),
	}
	c.cli.OnMessageCb = c.OnMessage
	c.cli.ChMsg = make(chan *network.Message, 1)
	c.chInput = make(chan *InputParam, 1)
	c.console.chInput = c.chInput
	//https://github.com/phuhao00/greatestworks-proto.git
	//github.com/phuhao00/greatestworks-proto
	return c
}

func (c *Client) Run() {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		fmt.Println(err)
	}

	go c.Write(conn)
	go c.Read(conn)

}

func (c *Client) Read(conn net.Conn) {

	for {
		//time.Sleep(1* time.Second)
		message, err := c.packer.UnPack(conn)
		//fmt.Println("message and err := ", message, err)
		if _, ok := err.(net.Error); err != nil && ok {
			fmt.Println("client err", err)
			continue
		}

		fmt.Println("resp message ", string(message.Data))
	}

}

func (c *Client) Write(conn net.Conn) {
	tick := time.NewTicker(time.Second)

	for {
		select {
		case <-tick.C:
			c.Send(conn, &Message{
				Id:   111,
				Data: []byte("-hello"),
			})
		}
	}
}

func (c *Client) Send(conn net.Conn, message *Message) {

	if err := conn.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
		return
	}

	bytes, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("network Send := ", string(bytes))
	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}

}
