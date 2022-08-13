package ziface


//定义一个服务接口

type IService interface {
	//启动服务器
	Start()

	//停止服务器
	Stop()

	//运行服务器
	Serve()

	//路由方法
	AddRouter(msgId uint32 , router IRouter)

	//获取service的连接
	GetConnMgr() IConnManager
}