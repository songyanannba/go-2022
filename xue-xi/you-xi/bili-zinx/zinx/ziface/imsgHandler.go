package ziface

type IMsgHandle interface {
	//调度执行router
	DoMsgHandler(request IRequest)

	//为消息添加具体的处理逻辑
	AddRouter(msgID uint32 , router IRouter)
}

