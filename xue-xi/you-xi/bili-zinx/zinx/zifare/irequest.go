package zifare

/**
	IRequest接口
	实际上 是把客户端的请求数据 包装到一个结构体中
 */
type IRequest interface {
	//当前的连接
	GetConnection() IConnection

	//请求的消息数据
	GetData() []byte
}