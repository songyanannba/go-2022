package znet

import (
	"bili-zinx/zinx/ziface"
	"errors"
	"fmt"
	"io"
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

	//告知 当前连接退出的channel
	ExitChan chan bool

	//该连接处理的方法
	MsgHandle ziface.IMsgHandle
}

//初始化连接的方法

func NewConnection(conn *net.TCPConn, connID uint32, msgHandle ziface.IMsgHandle) *Connection {
	c := &Connection{
		conn:     conn,
		ConnId:   connID,
		isClosed: false,
		MsgHandle:   msgHandle,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("reader Goroutine is running...")
	defer fmt.Println("connID = ", c.ConnId, "reader is exit , remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中

		//创建 拆包解包对象
		dp := NewDataPack()
		headData := make([]byte , dp.GetHeadLen())

		//读取客户端msg head 8个字节
		if _, err := io.ReadFull(c.GetTCPConnection(), headData) ; err != nil {
			fmt.Println("read msg head error", err)
			break
		}

		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack err" , err)
			break
		}

		//拆包 得到msgId msgDataLen 放到msg消息中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte , msg.GetMsgLen())
			//根据msgLen再次去读 data数据
			if _, err := io.ReadFull(c.GetTCPConnection(), data) ; err != nil {
				fmt.Println("read msg data error" , err)
				break
			}
		}
		msg.SetData(data)

		//当前连接的request
		req := Request{
			conn: c,
			msg: msg,
		}

		//c.Router.PreHandle(&req)
		//fmt.Println("ssss==" ,string(buf))
		//调用路由 执行方法
		go c.MsgHandle.DoMsgHandler(&req)

	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnId)
	//启动 从 当前连接的读业务
	go c.StartReader()

	//todo 启动当前连接的写业务
}

func (c *Connection) Stop() {
	fmt.Println("conn stop()... ConnID ", c.ConnId)
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

func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {

	return c.conn.RemoteAddr()
}

//提供一个sendMsg方法 ；封包 并发送给客户端

func (c *Connection) SendMsg(msgId uint32 ,data []byte) error {

	if c.isClosed == true {
		return errors.New("conn close when send msg...")
	}
	
	//封包
	dp := NewDataPack()
	
	message :=NewMessagePack(msgId ,data)

	binaryMsg, err := dp.Pack(message)
	if err != nil {
		fmt.Println("msg pack err")
		return err
	}

	if _, err := c.conn.Write(binaryMsg) ; err != nil {
		fmt.Println("write msg err msg id" , msgId , "error:" ,err)
		return err
	}

	return nil
}
