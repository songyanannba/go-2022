package ziface

type IMsgHandle interface {
	//调度执行router
	DoMsgHandler(request IRequest)

	//为消息添加具体的处理逻辑
	AddRouter(msgID uint32 , router IRouter)


	StartWorkerPool()

	//将消息发给任务队列处理
	SendMsgToTaskQueue(request IRequest)
}

