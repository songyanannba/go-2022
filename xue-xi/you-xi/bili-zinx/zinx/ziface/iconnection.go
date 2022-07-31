package ziface

import "net"

type IConnection interface {
	//启动
	Start()
	//停止
	Stop()
	//获取当前连接socket
	GetTCPConnection() *net.TCPConn

	//获取当前连接的id
	GetConnID() uint32

	//获取远程客户端的 tcp状态 IP port
	RemoteAddr() net.Addr

	//发送数据 将数据发送给远程客户端
	Send(data []byte) error
}

//定义一个处理业务的方法
type HandleFunc func(*net.TCPConn , []byte , int) error