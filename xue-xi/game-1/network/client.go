package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Client struct {
	Address string
	packer  NormalPacker
}

func NewClient(address string) *Client {
	return &Client{
		Address: address,
		packer: NormalPacker{
			Order: binary.BigEndian,
		},
	}
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
