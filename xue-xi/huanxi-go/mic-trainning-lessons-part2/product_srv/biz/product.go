package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-training-lessons-part2/proto/pb"
)

type ProductServer struct {
}

func (p ProductServer) ProductList(ctx context.Context, req *pb.ProductConditionReq) (*pb.ProductsRes, error) {
	panic("implement me")
}

func (p ProductServer) BatchGetProduct(ctx context.Context, req *pb.BatchProductIdReq) (*pb.ProductsRes, error) {
	panic("implement me")
}

func (p ProductServer) CreateProduct(ctx context.Context, item *pb.CreateProductItem) (*pb.ProductItemRes, error) {
	panic("implement me")
}

func (p ProductServer) DeleteProduct(ctx context.Context, item *pb.ProductDelItem) (*emptypb.Empty, error) {
	panic("implement me")
}

func (p ProductServer) UpdateProduct(ctx context.Context, item *pb.CreateProductItem) (*emptypb.Empty, error) {
	panic("implement me")
}

func (p ProductServer) GetProductDetail(ctx context.Context, res *pb.ProductItemRes) (*pb.ProductItemRes, error) {
	panic("implement me")
}


