package network

import (
	"fmt"
	"net"
)

type Server struct {
	Listener net.Listener
	Address  string
	NetWork  string
}

func NewService(Address, netWork string) *Server {
	return &Server{
		Listener: nil,
		Address:  Address,
		NetWork:  netWork,
	}
}

func (s *Server) Run() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", s.Address)
	if err != nil {
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return
	}

	s.Listener = tcpListener

	go func() {
		for {
			conn, err := s.Listener.Accept()

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("accept go deal ")
			go func() {
				session := NewSession(conn)
				session.Run()
			}()
		}
	}()

}

//接收包

//发送包
