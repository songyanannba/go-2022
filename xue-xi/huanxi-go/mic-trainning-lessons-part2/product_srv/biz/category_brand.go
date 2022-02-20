package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-training-lessons-part2/proto/pb"
)

func (p ProductServer) CategoryBrandList(ctx context.Context, req *pb.PagingReq) (*pb.CategoryBrandListRes, error) {
	panic("implement me")
}

func (p ProductServer) GetCategoryBrandList(ctx context.Context, req *pb.CategoryItemReq) (*pb.BrandItemRes, error) {
	panic("implement me")
}

func (p ProductServer) CreateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*pb.CategoryBrandRes, error) {
	panic("implement me")
}

func (p ProductServer) DeleteCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	panic("implement me")
}

func (p ProductServer) UpdateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	panic("implement me")
}
