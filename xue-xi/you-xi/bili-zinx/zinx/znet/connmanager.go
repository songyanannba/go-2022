package znet

import (
	"bili-zinx/zinx/ziface"
	"errors"
	"fmt"
	"sync"
)

type ConnManager struct {
	Connection map[uint32]ziface.IConnection
	connLock   sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		Connection: make(map[uint32]ziface.IConnection),
	}
}

//添加连接
func (connMgr *ConnManager) Add(conn ziface.IConnection) {

	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	connMgr.Connection[conn.GetConnID()] = conn
	fmt.Println("connecion add conn to connManager succ ,GetConnID = ", conn.GetConnID())
}

//删除连接
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {

	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	delete(connMgr.Connection, conn.GetConnID())

	fmt.Println("connecion Remove conn to connManager succ ,GetConnID = ", conn.GetConnID())

}

//获取当前连接
func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {

	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	if connection, ok := connMgr.Connection[connID]; ok {
		return connection, nil
	}
	return nil, errors.New("connection not found")

}

//得到当前的连接总数
func (connMgr *ConnManager) Len() int {
	return len(connMgr.Connection)
}

//清除并终止所有的连接
func (connMgr *ConnManager) ClearConn() {
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()

	for connID , conn :=  range connMgr.Connection {
		conn.Stop()
		delete(connMgr.Connection , connID)
	}

	fmt.Println("connecion ClearConn  connManager succ connLen = " ,connMgr.Len())
}
