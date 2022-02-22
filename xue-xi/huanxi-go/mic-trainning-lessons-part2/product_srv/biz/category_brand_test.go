package biz

import (
	"context"
	"fmt"
	"mic-training-lessons-part2/proto/pb"
	"testing"
)

func TestProductServer_CreateCategoryBrand(t *testing.T) {
	brand, err := client.CreateCategoryBrand(context.Background(), &pb.CategoryBrandReq{
		CategoryId: 39,
		BrandId:    2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(brand)
}

func TestProductServer_UpdateCategoryBrand(t *testing.T) {
	brand, err := client.UpdateCategoryBrand(context.Background(), &pb.CategoryBrandReq{
		Id:         1,
		BrandId:    3,
		CategoryId: 40,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(brand)
}

func TestProductServer_DeleteCategoryBrand(t *testing.T) {
	category, err := client.DeleteCategoryBrand(context.Background(), &pb.CategoryBrandReq{Id: 2})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(category)
}

func TestProductServer_CategoryBrandList(t *testing.T) {
	list, err := client.CategoryBrandList(context.Background(), &pb.PagingReq{
		PageNo:   1,
		PageSize: 5,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list)
}

func TestProductServer_GetCategoryBrandList(t *testing.T) {
	list, err := client.GetCategoryBrandList(context.Background(), &pb.CategoryItemReq{
		Id:               39,
		ParentCategoryId: 40,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list)
}
