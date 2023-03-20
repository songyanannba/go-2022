package biz

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-training-lessons-part2/proto/pb"
	"testing"
)


func TestProductServer_CreateAdvertise(t *testing.T) {
	for i := 0; i < 9; i++ {
		advertise, _ := client.CreateAdvertise(context.Background(), &pb.AdvertiseReq{
			Index: int32(i),
			Image: fmt.Sprintf("image-%d", i),
			Url:   fmt.Sprintf("url-%d", i),
		})
		fmt.Println(advertise.Id)
	}
}

func TestProductServer_AdvertiseList(t *testing.T) {
	adList, err := client.AdvertiseList(context.Background(), &emptypb.Empty{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(adList.Total)
	fmt.Println(adList.ItemList)
}

func TestProductServer_DeleteAdvertise(t *testing.T) {
	advertise, err := client.DeleteAdvertise(context.Background(), &pb.AdvertiseReq{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(advertise)
}

func TestProductServer_UpdateAdvertise(t *testing.T) {
	client.UpdateAdvertise(context.Background() , &pb.AdvertiseReq{
		Id:    2,
		Index: 2,
		Image: "hhh",
		Url:   "vvv",
	})
}