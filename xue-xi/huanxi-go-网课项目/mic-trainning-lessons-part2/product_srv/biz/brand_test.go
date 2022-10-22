package biz

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-training-lessons-part2/internal"
	"mic-training-lessons-part2/proto/pb"
	"testing"
)

var client pb.ProductServiceClient

func init() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.ProductSrvConfig.Host, internal.AppConf.ProductSrvConfig.Port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = pb.NewProductServiceClient(conn)
}

func TestProductServer_CreateBrand(t *testing.T) {
	brands := []string{
		"大十字", "横渡丽水", "原省状元", "xiangniiuyisai", "凯西小妞", "大黄飞机",
	}
	for _, item := range brands {
		res, err := client.CreateBrand(context.Background(), &pb.BrandItemReq{
			Name: item,
			Logo: "https://space.bilibili.com/375038855",
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res.Id)
	}
}

func TestProductServer_BrandList(t *testing.T) {
	res , err := client.BrandList(context.Background(),&pb.BrandPagingReq{
		PageNo: 2,
		PageSize: 5,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Total)
	for _ , ItemList := range res.ItemList {
		fmt.Println(ItemList)
	}

}

func TestProductServer_DeleteBrand(t *testing.T) {
	res, err := client.DeleteBrand(context.Background(), &pb.BrandItemReq{
		Id: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}


func TestProductServer_UpdateBrand(t *testing.T) {
	res, err := client.UpdateBrand(context.Background(), &pb.BrandItemReq{
		Id:   2,
		Name: "齐天大傻吧",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}