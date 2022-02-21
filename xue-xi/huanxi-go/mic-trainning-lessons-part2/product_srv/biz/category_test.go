package biz

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-training-lessons-part2/proto/pb"
	"testing"
)

func TestProductServer_CreateCategory(t *testing.T) {
	//第一级
	res, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "鲜肉",
		ParentCategoryId: 10,
		Level:            2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

	//第二级
	res2, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛肉",
		ParentCategoryId: res.Id,
		Level:            3,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res2)

	//第三级
	res3, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛排",
		ParentCategoryId: res2.Id,
		Level:            4,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res3)
}

func TestProductServer_DeleteCategory(t *testing.T) {

	client.DeleteCategory(context.Background(), &pb.CategoryDelReq{
		Id: 42,
	})
}

func TestProductServer_UpdateCategory(t *testing.T) {
	category, err := client.UpdateCategory(context.Background(), &pb.CategoryItemReq{
		Id:   41,
		Name: "建起吗",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(category)
}

func TestProductServer_GetAllCategoryList(t *testing.T) {
	list, err := client.GetAllCategoryList(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list)
}


func TestProductServer_GetSubCategory(t *testing.T) {
	category, err := client.GetSubCategory(context.Background(), &pb.CategoriesReq{
		Id:    39,
		Level: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(category)
}