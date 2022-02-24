package biz

import (
	"context"
	"fmt"
	"mic-training-lessons-part2/proto/pb"
	"testing"
)

func TestProductServer_CreateProduct(t *testing.T) {
	for i := 0; i < 8; i++ {
		product, err := client.CreateProduct(context.Background(), &pb.CreateProductItem{
			Name:        fmt.Sprintf("黄金牛排%d", i),
			Sn:          "1234234",
			Price:       100.00,
			RealPrice:   200.00,
			ShortDesc:   "哈哈哈",
			ProductDesc: "收拾收拾",
			Images:      nil,
			DescImages:  nil,
			CoverImage:  "www",
			IsNew:       true,
			IsPop:       true,
			Selling:     true,
			BrandId:     3,
			FavNum:      666,
			SoldNum:     888,
			CategoryId:  40,
			IsShipFree:  true,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(product)
	}
}

func TestProductServer_UpdateProduct(t *testing.T) {
	client.UpdateProduct(context.Background(), &pb.CreateProductItem{
		Id:         8,
		Name:       "hjkdsgfs",
		CategoryId: 41,
		BrandId:    6,
	})
}

func TestProductServer_DeleteProduct(t *testing.T) {
	client.DeleteProduct(context.Background(), &pb.ProductDelItem{
		Id: 7,
	})
}

func TestProductServer_BatchGetProduct(t *testing.T) {
	ids := []int32{2, 3, 4}
	product, _ := client.BatchGetProduct(context.Background(), &pb.BatchProductIdReq{Ids: ids})
	fmt.Println(product)
}

func TestProductServer_ProductList(t *testing.T) {
	list, _ := client.ProductList(context.Background(), &pb.ProductConditionReq{
		PageSize: 10,
		PageNo:   1,
	})
	fmt.Println(list)
}