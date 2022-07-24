package znet

import "bili-zinx/zinx/zifare"

type Request struct {
	//和客户端的连接
	conn zifare.IConnection

	//客户端的请求数据
	data []byte
}

func (r *Request)GetConnection() zifare.IConnection {
	return r.conn
}

func(r *Request)GetData() []byte {
	return r.data
}