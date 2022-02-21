package biz

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-training-lessons-part2/custom_error"
	"mic-training-lessons-part2/internal"
	"mic-training-lessons-part2/model"
	"mic-training-lessons-part2/proto/pb"
)

func (p ProductServer) CreateCategory(ctx context.Context, req *pb.CategoryItemReq) (*pb.CategoryItemRes, error) {
	category := model.Category{}
	//todo 参数判断
	category.Name = req.Name
	category.Level = req.Level
	if category.Level > 1 {
		category.ParentCategoryID = req.ParentCategoryId
	}
	internal.DB.Save(&category)
	res := ConvertCategoryModel2Pb(category)
	return res, nil
}

func (p ProductServer) GetAllCategoryList(ctx context.Context, empty *emptypb.Empty) (*pb.CategoriesRes, error) {
	var categoryList []model.Category
	internal.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categoryList)

	var res pb.CategoriesRes

	/*var items []*pb.CategoryItemRes
	for _, item := range categoryList {
		items = append(items, ConvertCategoryModel2Pb(item))
	}*/

	b, err := json.Marshal(categoryList)
	if err != nil {
		return nil, errors.New(custom_error.MarshalFails)
	}
	//res.InfoResList = items
	res.CategoryJsonFormat = string(b)

	return &res, nil
}

func (p ProductServer) GetSubCategory(ctx context.Context, req *pb.CategoriesReq) (*pb.SubCategoriesRes, error) {
	var category model.Category
	var subItemList []*pb.CategoryItemRes
	var res pb.SubCategoriesRes

	r := internal.DB.Find(&category, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}
	pre := "SubCategory"
	if category.Level == 1 {
		pre = "SubCategory.SubCategory"
	}
	var subCategoryList []model.Category
	internal.DB.Where(&model.Category{ParentCategoryID: req.Id}).Preload(pre).Find(&subCategoryList)

	for _, item := range subCategoryList {
		subItemList = append(subItemList, ConvertCategoryModel2Pb(item))
	}
	b, err := json.Marshal(subItemList)
	if err != nil {
		return nil, errors.New(custom_error.MarshalFails)
	}

	res.SubCategoryList = subItemList
	res.CategoryJsonFormat = string(b)
	return &res, nil
}

func (p ProductServer) DeleteCategory(ctx context.Context, req *pb.CategoryDelReq) (*emptypb.Empty, error) {
	internal.DB.Delete(&model.Category{}, req.Id)
	return &emptypb.Empty{}, nil
}

func (p ProductServer) UpdateCategory(ctx context.Context, req *pb.CategoryItemReq) (*emptypb.Empty, error) {
	var category model.Category

	r := internal.DB.Find(&category, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategoryId > 0 {
		category.ParentCategoryID = req.ParentCategoryId
	}
	if req.Level > 0 {
		category.Level = req.Level
	}

	internal.DB.Save(&category)
	return &emptypb.Empty{}, nil
}

func ConvertCategoryModel2Pb(item model.Category) *pb.CategoryItemRes {
	it := &pb.CategoryItemRes{
		Id:    item.ID,
		Name:  item.Name,
		Level: item.Level,
	}
	if item.Level > 1 {
		it.ParentCategoryId = item.ParentCategoryID
	}
	return it
}
