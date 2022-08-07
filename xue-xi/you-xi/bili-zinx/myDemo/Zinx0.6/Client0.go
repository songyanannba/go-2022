package main

import (
	"bili-zinx/zinx/znet"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("client start ...")

	time.Sleep(1 * time.Second)
	//连接远程服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		panic(err)
	}

	//写数据
	for {

		dp := znet.NewDataPack()
		binaryMsg , err := dp.Pack(znet.NewMessagePack(0 , []byte("ZinxV0.6 client Test Message")))

		if err != nil {
			fmt.Println("pack error:" , err)
		}

		_, err = conn.Write(binaryMsg)
		if err != nil {
			fmt.Println(" conn Write error:" , err)
			return
		}

		//服务器 回复一个message数据

		//先读区流中的 head部分 ， 得到ID 和dataLen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead) ; err != nil {
			fmt.Println("read head error" , err)
			break
		}

		msgHead, err := dp.UnPack(binaryHead)
		if err != nil {
			fmt.Println("client unpack err")
		}

		if msgHead.GetMsgLen() > 0 {
			//再根据DataLen进行 二次读取  将data读取出来
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte , msg.GetMsgLen())

			if _ , err = io.ReadFull(conn , msg.Data) ; err != nil {
				fmt.Println("read msg data err")
				return
			}

			fmt.Println("---> Recv Server Msg :ID = " , msg.Id , "len = " , msg.DataLen , "data = " , string(msg.Data) )
		}



		//cpu 阻塞
		time.Sleep(1 *time.Second)

	}

}

