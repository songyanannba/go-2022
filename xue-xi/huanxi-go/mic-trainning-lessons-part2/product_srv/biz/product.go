package biz

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-training-lessons-part2/custom_error"
	"mic-training-lessons-part2/internal"
	"mic-training-lessons-part2/model"
	"mic-training-lessons-part2/proto/pb"
)

type ProductServer struct {
}

func (p ProductServer) ProductList(ctx context.Context, req *pb.ProductConditionReq) (*pb.ProductsRes, error) {
	iDb := internal.DB.Model(model.Product{})
	var productList []model.Product
	var itemList []*pb.ProductItemRes
	var res *pb.ProductsRes

	if req.IsPop {
		iDb = iDb.Where("is_pop=?", req.IsPop)
	}

	if req.IsNew {
		iDb = iDb.Where("is_new=?", req.IsNew)
	}

	if req.BrandId > 0 {
		iDb = iDb.Where("brand_id=?", req.BrandId)
	}

	if req.KeyWord != "" {
		iDb = iDb.Where("key_word like ?", "%"+req.KeyWord+"%")
	}

	if req.MinPrice > 0 {
		iDb = iDb.Where("min_price > ?", req.MinPrice)
	}

	if req.MaxPrice > 0 {
		iDb = iDb.Where("man_price > ?", req.MaxPrice)
	}

	if req.CategoryId > 0 {
		var category model.Category
		r := internal.DB.First(&category, req.CategoryId)
		if r.RowsAffected < 1 {
			return nil, errors.New(custom_error.CategoryNotExist)
		}
		var q string
		if category.Level == 1 {
			q = fmt.Sprintf("select id from category where parent_category_id in (select id from category where id=%d)", req.CategoryId)
		} else if category.Level == 2 {
			q = fmt.Sprintf("select id from category where parent_category_id=%d", req.CategoryId)
		} else if category.Level == 3 {
			q = fmt.Sprintf("select id from category WHERE id=%d", req.CategoryId)
		}
		iDb = iDb.Where("category_id in %s", q)
	}

	var count int64
	iDb.Count(&count)
	fmt.Println(count)
	iDb.Joins("Category").Joins("Brand").Scopes(internal.MyPaging(int(req.PageNo), int(req.PageSize))).Find(&productList)

	for _, item := range productList {
		res2 := ConvertProductModel2Pb(item)
		itemList = append(itemList, res2)
	}
	res.Total = int32(count)
	res.ItemList = itemList
	return res, nil
}

func (p ProductServer) BatchGetProduct(ctx context.Context, req *pb.BatchProductIdReq) (*pb.ProductsRes, error) {
	var productList []model.Product
	var res pb.ProductsRes

	r := internal.DB.Find(&productList, req.Ids)

	res.Total = int32(r.RowsAffected)

	for _, item := range productList {
		pro := ConvertProductModel2Pb(item)
		res.ItemList = append(res.ItemList, pro)
	}

	return &res, nil
}

func (p ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductItem) (*pb.ProductItemRes, error) {

	var category model.Category
	var brand model.Brand
	var res *pb.ProductItemRes

	r := internal.DB.First(&category, req.CategoryId)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}

	b := internal.DB.First(&brand, req.BrandId)
	if b.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}

	item := ConvertReq2Model(req, category, brand)

	internal.DB.Save(&item)
	res = ConvertProductModel2Pb(item)
	return res, nil
}

func (p ProductServer) DeleteProduct(ctx context.Context, req *pb.ProductDelItem) (*emptypb.Empty, error) {
	r := internal.DB.Delete(&model.Product{}, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.DelProductFail)
	}
	return &emptypb.Empty{}, nil
}

func (p ProductServer) UpdateProduct(ctx context.Context, req *pb.CreateProductItem) (*emptypb.Empty, error) {

	var product model.Product
	var category model.Category
	var brand model.Brand

	r1 := internal.DB.First(&product, req.Id)
	if r1.RowsAffected < 1 {
		return nil, errors.New(custom_error.ProductNotExits)
	}

	r := internal.DB.First(&category, req.CategoryId)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}

	b := internal.DB.First(&brand, req.BrandId)
	if b.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}

	/*req2Model := ConvertReq2Model(req, category, brand)
	internal.DB.Save(req2Model)*/

	return &emptypb.Empty{}, nil

}

func (p ProductServer) GetProductDetail(ctx context.Context, req *pb.ProductItemRes) (*pb.ProductItemRes, error) {
	var product model.Product
	var res *pb.ProductItemRes

	r := internal.DB.First(&product, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.ProductNotExits)
	}
	res = ConvertProductModel2Pb(product)

	return res, nil
}

func ConvertReq2Model(req *pb.CreateProductItem, category model.Category, brand model.Brand) model.Product {
	p := model.Product{
		CategoryID: req.CategoryId,
		BrandID:    req.BrandId,
		Selling:    req.Selling,
		IsShipFree: req.IsShipFree,
		IsPop:      req.IsPop,
		IsNew:      req.IsNew,
		Name:       req.Name,
		SN:         req.Sn,
		FavNum:     req.FavNum,
		SoldNum:    req.SoldNum,
		Price:      req.Price,
		RealPrice:  req.RealPrice,
		ShorDesc:   req.ShortDesc,
		Images:     req.Images,
		DescImages: req.DescImages,
		CoverImage: req.CoverImage,
	}
	return p
}

func ConvertProductModel2Pb(p model.Product) *pb.ProductItemRes {
	pb := &pb.ProductItemRes{
		Id:         p.ID,
		CategoryId: p.CategoryID,
		Name:       p.Name,
		Sn:         p.SN,
		FavNum:     p.FavNum,
		Price:      p.Price,
		RealPrice:  p.RealPrice,
		ShortDesc:  p.ShorDesc,
		Images:     p.Images,
		DescImages: p.DescImages,
		CaverImage: p.CoverImage,
		IsNew:      p.IsNew,
		IsPop:      p.IsPop,
		Selling:    p.Selling,
		Category: &pb.CategoryItemRes{
			Id:   p.Category.ID,
			Name: p.Category.Name,
		},
		Brand: &pb.BrandItemRes{
			Id:   p.Brand.ID,
			Name: p.Brand.Name,
			Logo: p.Brand.Logo,
		},
	}
	return pb
}
