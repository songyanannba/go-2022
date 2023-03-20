package ziface

type IDataPack interface {

	GetHeadLen() uint32

	//封包
	Pack(msg IMessage) ([]byte , error)

	//拆包
	UnPack([]byte) (IMessage , error)

}
