package znet

import (
	"bili-zinx/zinx/zifare"
	"fmt"
	"net"
)

/**
连接模块
 */
type Connection struct {
	//当前连接的 tcp套接字
	conn *net.TCPConn

	//链接的id
	ConnId uint32

	//当前的连接状态
	isClosed bool

	//当前连接的绑定业务方法api
	handleAPI zifare.HandleFunc

	//告知 当前连接退出的channel
	ExitChan chan bool
}

//初始化连接的方法

func NewConnection(conn *net.TCPConn , connID uint32 , callbackApi zifare.HandleFunc )  *Connection {
	c:= &Connection{
		conn:      conn,
		ConnId:    connID,
		isClosed:  false,
		handleAPI: callbackApi,
		ExitChan:  make(chan bool , 1),
	}
	return c
}

func (c *Connection)StartReader()  {
	fmt.Println("reader Goroutine is running...")
	defer fmt.Println("connID = " ,c.ConnId , "reader is exit , remote addr is " , c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		buf := make([]byte ,512)
		cnt ,err := c.conn.Read(buf)
		if err != nil {
			fmt.Println("reve buf err" , err)
			continue
		}
		//调用 当前绑定的handleApi
		if err := c.handleAPI(c.conn ,buf ,cnt) ;err != nil {
			fmt.Println("connID " , c.ConnId)
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnId)
	//启动 从 当前连接的读业务
	go c.StartReader()


	//todo 启动当前连接的写业务
}

func (c *Connection)Stop() {
	fmt.Println("conn stop()... ConnID " ,c.ConnId)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//关闭连接
	c.conn.Close()

	//回收资源
	close(c.ExitChan)
}


func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.conn
}


func (c *Connection) GetConnID () uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {

	return c.conn.RemoteAddr()
}

func (c *Connection) Send (data []byte) error {

	return nil
}