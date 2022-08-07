package znet

import (
	"bili-zinx/utils"
	"bili-zinx/zinx/ziface"
	"fmt"
)

type MsgHandle struct {
	//存放msgId 对应的方法
	Apis map[uint32]ziface.IRouter

	//消息队列
	TaskQueue []chan ziface.IRequest

	//worker工作池数量
	WorkerPoolSize uint32
}



func NewMsgHandle() *MsgHandle{
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,//配置文件
		TaskQueue: make([]chan ziface.IRequest , utils.GlobalObject.WorkerPoolSize),
	}
}


//调度执行router
func(mh *MsgHandle)DoMsgHandler(request ziface.IRequest) {

	//从request 里面找到msgID
	handler , ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = " , request.GetMsgID() , "is not fount need reqister")
	}

	//根据MsgID 调度对应的router业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

//为消息添加具体的处理逻辑
func(mh *MsgHandle)AddRouter(msgID uint32 , router ziface.IRouter) {

	if _, ok := mh.Apis[msgID] ; ok {
		//已经存在
		panic("router msg id not empty")
	}

	mh.Apis[msgID] = router
	fmt.Println("添加路由成功 msgID = " ,msgID)
}


//启动一个worker工作池
func (mh *MsgHandle) StartWorkerPool() {
	fmt.Println("StartWorkerPool")
	for i := 0 ; i < int(mh.WorkerPoolSize) ; i++ {
		//开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest , utils.GlobalObject.MaxWorkerTaskLen)
		//启动worker
		go mh.StartOneWorker(i ,mh.TaskQueue[i])
	}

}

//工作流程
func (mh *MsgHandle) StartOneWorker(workerId int , taskQueue chan ziface.IRequest) {
	fmt.Println("worker ID = " ,workerId ,"is started ...")
	for {
		select {
		case request := <- taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func(mh *MsgHandle)SendMsgToTaskQueue(request ziface.IRequest) {

	//将消息分配给不同的worker
	// 根据客户端的连接ID分配；平均分配
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnId = " , request.GetConnection().GetConnID(),
		"request MsgId = " ,request.GetMsgID() ,
		"to workerId =" , workerID)

	//将消息发给对应的worker的taskQueue
	mh.TaskQueue[workerID] <-request
}
