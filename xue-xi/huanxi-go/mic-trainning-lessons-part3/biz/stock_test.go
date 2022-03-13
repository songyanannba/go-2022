package biz

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-trainning-lessons-part3/internal"
	"mic-trainning-lessons-part3/proro/pb"
	"testing"
)

var client pb.StockServiceClient

func init() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.StockSrvConfig.Host,
		internal.AppConf.StockSrvConfig.Port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("init err")
	}
	client = pb.NewStockServiceClient(conn)
}

func TestStockServer_SetStock(t *testing.T) {
	client.SetStock(context.Background(), &pb.ProductStockItem{
		ProductId: 1,
		Num:       2,
	})
}

func TestStockServer_StockDetail(t *testing.T) {
	client.StockDetail(context.Background(), &pb.ProductStockItem{
		ProductId: 1,
	})
}

func TestStockServer_Sell(t *testing.T) {
	item := &pb.ProductStockItem{
		ProductId: 1,
		Num:       1,
	}
	var itemList []*pb.ProductStockItem

	itemList = append(itemList, item)
	sellItem := &pb.SellItem{
		StockItemList: itemList,
	}

	sell, err := client.Sell(context.Background(), sellItem)
	if err != nil {
		t.Fatal("sell err")
	}
	fmt.Println(sell)
}

func TestStockServer_BackStock(t *testing.T) {

	item := &pb.ProductStockItem{
		ProductId: 1,
		Num:       1,
	}
	var itemList []*pb.ProductStockItem

	itemList = append(itemList, item)
	sellItem := &pb.SellItem{
		StockItemList: itemList,
	}
	client.BackStock(context.Background(), sellItem)
}
