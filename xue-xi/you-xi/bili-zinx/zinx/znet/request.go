package znet

import "bili-zinx/zinx/ziface"

type Request struct {
	//和客户端的连接
	conn ziface.IConnection

	//客户端的请求数据
	data []byte
}

func (r *Request)GetConnection() ziface.IConnection {
	return r.conn
}

func(r *Request)GetData() []byte {
	return r.data
}