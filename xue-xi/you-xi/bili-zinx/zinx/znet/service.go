package znet

import (
	"bili-zinx/utils"
	"bili-zinx/zinx/ziface"
	"fmt"
	"net"
)

type Service struct {
	//服务器名字
	Name string
	//服务器绑定的IP版本
	IPVersion string
	//服务器监听的ip
	IP string
	//服务器监听的端口
	Port int

	//当前的消息管理模块
	MsgHandle ziface.IMsgHandle

	ConnMgr ziface.IConnManager

}


func (s *Service) Start() {
	fmt.Println("[Zinx start] name is " ,utils.GlobalObject.Name)
	fmt.Println("[Zinx start] Listener is " ,utils.GlobalObject.Host ,utils.GlobalObject.Port)
	fmt.Println("[Zinx start] Version is " ,utils.GlobalObject.Version)

	fmt.Printf("[start] Service Listenner at IP :%s ,Port %d is starting \n" ,s.IP ,s.Port)

	go func() {
		//0 开启消息队列和工作池
		s.MsgHandle.StartWorkerPool()

		//1 获取tcp 的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: " ,err)
			return
		}

		//2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp err: " ,err)
			return
		}

		fmt.Println("start Zinx service succ" ,s.Name )

		var cid uint32
		cid = 0

		//3 阻塞等待客户端连接 处理客户端连接业务
		//循环获取客户端的信息
		fmt.Println("cid === " ,cid)
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("conn err " ,err)
				continue
			}

			//最大连接的判断
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				//todo 给客户端相应一个最大连接的错误包
				fmt.Println("==>conn is too many" ,s.ConnMgr.Len())
				conn.Close()
				continue
			}

			dealConn := NewConnection(s, conn, cid, s.MsgHandle)
			cid++

			go dealConn.Start()

		}
	}()

}



func (s *Service)Stop() {
	//将一些服务器状态停止 或者回收
	fmt.Println("[STOP] zinx service stop")
	s.ConnMgr.ClearConn()
}

func (s *Service)Serve() {
	//启动服务
	s.Start()

	//做一些额外的工作

	//阻塞
	select {}
}

func (s *Service)AddRouter(msgId uint32 ,router ziface.IRouter) {
	//将一些服务器状态停止 或者回收
	s.MsgHandle.AddRouter(msgId , router)
	fmt.Println("add router succ...")
}

func(s *Service)GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

/**
初始化Service的方法
 */
func NewService () ziface.IService {

	s := &Service{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.Port,
		MsgHandle: NewMsgHandle(),
		ConnMgr: NewConnManager(),
	}

	return s
}