package biz

import (
	"context"
	"fmt"
	"mic-training-lessons-part2/proto/pb"
	"testing"
)

func TestProductServer_CreateCategory(t *testing.T) {
	//第一级
	res, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "鲜肉",
		ParentCategoryId: 10,
		Level:            1,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

	//第二级
	res2, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛肉",
		ParentCategoryId: res.Id,
		Level:            2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res2)

	//第三级
	res3, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛排",
		ParentCategoryId: res2.Id,
		Level:            2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res3)

}
