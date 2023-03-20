package main

import (
	"fmt"
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
		//cpu 阻塞
		time.Sleep(1 *time.Second)

		_, err := conn.Write([]byte("hello Zinx V0.3..."))
		if err != nil {
			fmt.Println("write err" , err)
			return
		}

		fmt.Println("Write succ and start read ...")
		buf := make([]byte ,512)
		cnt ,err := conn.Read(buf)

		fmt.Println("cnt err==")
		if err != nil {
			fmt.Println("read buf error" ,err)
			return
		}

		fmt.Printf("service call back :%s ,cnt = %d \n" , buf ,cnt)


	}

}

