package znet

import (
	"bili-zinx/zinx/zifare"
	"errors"
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
}



func (s *Service) Start() {
	fmt.Printf("[start] Service Listenner at IP :%s ,Port %d is starting \n" ,s.IP ,s.Port)

	go func() {
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

			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			go dealConn.Start()
			//原生
			/*go func() {
				for  {
					buf := make([]byte ,512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv read err " ,err)
						continue
					}

					fmt.Printf("recv client buf %s ,cnt %d\n" , buf ,cnt)
					//回显
					if _, err := conn.Write(buf[0:cnt]) ; err != nil {
						fmt.Println("write back buf err " , err)
						continue
					}
				}

			}()*/

		}
	}()

}

//定义客户端绑定的api
func CallBackToClient( conn *net.TCPConn , data []byte , cnt int) error {
	fmt.Println("call back client...")
	if _, err := conn.Write(data[:cnt]) ; err != nil {
		fmt.Println("write back buf err" , err)
		return errors.New("CallBackToClient err")
	}
	return nil
}

func (s *Service)Stop() {
	//将一些服务器状态停止 或者回收
}

func (s *Service)Serve() {
	//启动服务
	s.Start()

	//做一些额外的工作

	//阻塞
	select {}
}

/**
初始化Service的方法
 */
func NewService (Name string) zifare.IService {

	s := &Service{
		Name:      "name",
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}