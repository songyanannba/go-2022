package mian

import (
	"fmt"
	"mic-basic-lessons/week6/6-5/proto/pb"
)

func main() {
	req := pb.UpperRequest{
		Id:   1,
		Name: "面向加薪",
	}

	platform := pb.Platfrom{
		Name:     "B站",
		FansCont: "100",
	}

	resp := pb.UpperResponse{
		Name: "xjj哈哈",
		P:    &platform,
	}
	fmt.Println(req)
	fmt.Println(resp)

}
