package ziface

/**
	消息模块
 */

type IMessage interface {
	//获取
	GetMsgId() uint32

	GetMsgLen() uint32

	GetData() []byte


	//设置
	SetMsgId( uint32)

	SetData([]byte)

	SetDataLen(uint32)
}