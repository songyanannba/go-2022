package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part3/proro/pb"
)

type StockServer struct {
}

func (s StockServer) SetStock(ctx context.Context, req *pb.ProductStockItem) (*emptypb.Empty, error) {

	panic(nil)
}

func (s StockServer) StockDetail(ctx context.Context, req *pb.ProductStockItem) (*pb.ProductStockItem, error) {
	panic(nil)
}

func (s StockServer) Sell(ctx context.Context, req *pb.SellItem) (*emptypb.Empty, error) {
	panic(nil)
}

func (s StockServer) BackStock(ctx context.Context, req *pb.SellItem) (*emptypb.Empty, error) {
	panic(nil)
}
