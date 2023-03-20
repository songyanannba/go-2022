package main

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"time"
)

func main() {
	t1 := pb.TodoRequest{
		Todo: "面向加薪学习",
		Week: pb.Week_Monday,
	}

	fmt.Println(t1)
	fmt.Println(pb.Week_Friday)

	m1 := map[string]string{"go": "Go语言极简一本通"}
	t1.BookMap = m1
	fmt.Println(t1)

	t1.DoneTime = timestamppb.New(time.Now())

	fmt.Println(t1)

}
