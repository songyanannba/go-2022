package biz

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-training-lessons-part2/custom_error"
	"mic-training-lessons-part2/internal"
	"mic-training-lessons-part2/model"
	"mic-training-lessons-part2/proto/pb"
)

func (p ProductServer) CreateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*pb.CategoryBrandRes, error) {
	var res pb.CategoryBrandRes
	var item model.ProductCategoryBrand
	var category model.Category
	var brand model.Brand

	//分类判断
	r := internal.DB.Find(&category, req.CategoryId)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}
	//品牌判断
	b := internal.DB.Find(&brand, req.BrandId)
	if b.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}
	//是否已经存在关系
	item.CategoryID = req.CategoryId
	item.BrandID = req.BrandId

	internal.DB.Save(&item)

	res.Id = item.ID
	return &res, nil

}

func (p ProductServer) DeleteCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	r := internal.DB.Delete(&model.ProductCategoryBrand{}, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.DelProductCategoryBrandFail)
	}
	return &emptypb.Empty{}, nil
}

func (p ProductServer) UpdateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	var cateBrand model.ProductCategoryBrand
	cateBrand.ID = req.Id
	cateBrand.CategoryID = req.CategoryId
	cateBrand.BrandID = req.BrandId

	//分类判断
/*	r := internal.DB.Find(&cateBrand, req.CategoryId)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}*/
	//品牌判断
/*	b := internal.DB.Find(&cateBrand, req.BrandId)
	if b.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}*/

	internal.DB.Save(&cateBrand)
	return &emptypb.Empty{}, nil
}

func (p ProductServer) CategoryBrandList(ctx context.Context, req *pb.PagingReq) (*pb.CategoryBrandListRes, error) {

	var count int64
	var items []model.ProductCategoryBrand
	var res pb.CategoryBrandListRes
	var resList []*pb.CategoryBrandRes

	internal.DB.Model(&model.ProductCategoryBrand{}).Count(&count)
	internal.DB.Preload("Category").Preload("Brand").Scopes(internal.MyPaging(int(req.PageNo), int(req.PageSize))).Find(&items)
	for _, item := range items {
		pcb := ConvertProductCategoryBrand2Pb(item)
		resList = append(resList, pcb)
	}
	res.ItemList = resList
	res.Total = int32(count)

	return &res, nil

}

func (p ProductServer) GetCategoryBrandList(ctx context.Context, req *pb.CategoryItemReq) (*pb.BrandRes, error) {

	var res pb.BrandRes
	var category model.Category
	var itemList []model.ProductCategoryBrand
	var itemListRes []*pb.BrandItemRes

	r := internal.DB.Find(&category, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExist)
	}

	r = internal.DB.Preload("Brand").Where(&model.ProductCategoryBrand{CategoryID: req.ParentCategoryId}).Find(&itemList)

	if r.RowsAffected > 0 {
		res.Total = int32(r.RowsAffected)
	}

	for _, item := range itemList {
		itemListRes = append(itemListRes, &pb.BrandItemRes{
			Id:   item.Brand.ID,
			Name: item.Brand.Name,
			Logo: item.Brand.Logo,
		})
	}
	res.ItemList = itemListRes
	return &res, nil
}

func ConvertProductCategoryBrand2Pb(pcb model.ProductCategoryBrand) *pb.CategoryBrandRes {
	cd := &pb.CategoryBrandRes{
		Id: pcb.ID,
		Brand: &pb.BrandItemRes{
			Id:   pcb.Brand.ID,
			Name: pcb.Brand.Name,
			Logo: pcb.Brand.Logo,
		},
		Category: &pb.CategoryItemRes{
			Id:               pcb.Category.ID,
			Name:             pcb.Category.Name,
			ParentCategoryId: pcb.Category.ParentCategoryID,
			Level:            pcb.Category.Level,
		},
	}
	return cd
}
