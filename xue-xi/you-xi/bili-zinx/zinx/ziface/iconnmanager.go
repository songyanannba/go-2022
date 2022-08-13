package ziface


type IConnManager interface {

	//添加连接
	Add(conn IConnection)

	//删除连接
	Remove(conn IConnection)

	//获取当前连接
	Get(connID uint32)(IConnection , error)

	//得到当前的连接总数
	Len() int

	//清除并终止所有的连接
	ClearConn()

}