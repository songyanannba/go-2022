package network

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Session struct {
	Conn    net.Conn
	Packer  *NormalPacker
	chWrite chan *Message
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		Conn:    conn,
		Packer:  NewNormalPacker(binary.BigEndian),
		chWrite: make(chan *Message, 10),
	}
}

func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

func (s *Session) Read() {
	//time.Sleep(2 * time.Second)
	/*if err := s.Conn.SetReadDeadline(time.Now().Add(time.Second )); err != nil {
		fmt.Println( "session Read ",err)
	}*/
	fmt.Println("session Read for...")
	for {
		message, err := s.Packer.UnPack(s.Conn)
		//fmt.Println("message and err := " , message , err)
		if err != nil {
			fmt.Println("session Read for", err)
			return
		}

		fmt.Println("net work session Read", string(message.Data))

		s.chWrite <- &Message{
			Id:   999,
			Data: []byte("hi syn"),
		}
	}

}

func (s *Session) Write() {

	/*if err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
		fmt.Println(err)
		return
	}*/

	for {
		select {
		case msg := <-s.chWrite:
			s.Send(msg)
		}
	}

}

func (s *Session) Send(message *Message) {

	fmt.Println("session send...")
	pack, err := s.Packer.Pack(message)
	if err != nil {
		fmt.Println(err)
	}

	_, err = s.Conn.Write(pack)
	if err != nil {
		fmt.Println(err)
	}

}
