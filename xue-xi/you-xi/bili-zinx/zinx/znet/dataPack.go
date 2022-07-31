package znet

import (
	"bili-zinx/utils"
	"bili-zinx/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

/**
	封包 拆包 模块
 */

type DataPack struct {

}

func NewDataPack() *DataPack{
	return &DataPack{}
}


func(dp *DataPack) GetHeadLen() uint32 {
	// dataLen uint32 (4字节)
	// id uint32 (4字节)
	return 8
}

//封包
func(dp *DataPack) Pack(msg ziface.IMessage) ([]byte , error) {
	dataBuff := bytes.NewBuffer([]byte{})
	//将数据写入dataBuff中
	//将长度 写入
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil ,err
	}
	//将消息ID 写入
	if err := binary.Write(dataBuff , binary.LittleEndian , msg.GetMsgId()) ; err != nil {
		return nil ,err
	}
	//将data 写入
	if err := binary.Write(dataBuff , binary.LittleEndian , msg.GetData()) ; err != nil {
		return nil ,err
	}

	return dataBuff.Bytes() , nil
}

//拆包
func(dp *DataPack)  UnPack(binData []byte) (ziface.IMessage , error) {
	//先读区头部 在读内容

	//创建一个读取器
	dataBuff := bytes.NewReader(binData)

	msg := &Message{}

	//读取dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen) ; err != nil {
		return nil , err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data ")
	}

	//读取消息ID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id) ; err != nil {
		return nil , err
	}


	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Data) ; err != nil {
		return nil , err
	}

	return msg , nil
}