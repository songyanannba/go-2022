package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {

	/**
		模拟服务器
	 */
	 listenner , err := net.Listen("tcp" , "127.0.0.1:7777")
	 if err != nil {
		fmt.Println("service listen err: " , err)
		return
	}
	
	go func() {
		//从客户端读取数据 拆包处理
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("service accept error ", err)
			}

			go func(conn net.Conn) {
				//处理客户端的请求
				//拆包的过程
				dp := NewDataPack()
				for {
					//读取head
					headData := make([]byte ,dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head err" , err)
						return
					}

					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("server UnPack err" , err)
						return
					}

					if msgHead.GetMsgLen() > 0 {
						//有数据 再次读取
						msg := msgHead.(*Message)
						msg.Data = make([]byte , msg.GetMsgLen())

						if _ , err = io.ReadFull(conn ,msg.Data) ; err != nil {
							fmt.Println("server UnPack data err" , err)
							return
						}
						fmt.Println("-->rece Msg succ len , id ,data = " ,msg.DataLen , msg.Id , string(msg.Data))
					}

				}

			}(conn)
		}
	}()


	/**
		模拟客户端
	 */
	conn , err := net.Dial("tcp" , "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:" , err)
		return
	}

	//创建封包对象
	dp := NewDataPack()

	//封装第一个包
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z' , 'i' , 'n' , 'x' },
	}

	sendData1 , err := dp.Pack(msg1)

	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'n' , 'h' ,'z' , 'i' , 'n' , 'x' ,'!'},
	}

	sendData2 , err := dp.Pack(msg2)

	sendData := append(sendData1 , sendData2...)

	conn.Write(sendData)

	//客户端阻塞
	select {}
}