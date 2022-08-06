package znet

import (
	"bili-zinx/zinx/ziface"
	"fmt"
)

type MsgHandle struct {

	Apis map[uint32]ziface.IRouter
}


func NewMsgHandle() *MsgHandle{
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
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