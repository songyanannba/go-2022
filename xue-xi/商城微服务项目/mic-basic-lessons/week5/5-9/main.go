package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"mic-basic-lessons/week5/5-9/ptoto/pb"
)

func main() {
	req := pb.BoolRequest{Name: "go语言极简一本通"}
	b ,err := proto.Marshal(&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))

}
