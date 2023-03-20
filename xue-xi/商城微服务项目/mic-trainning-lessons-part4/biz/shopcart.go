package biz

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part4/custom_error"
	"mic-trainning-lessons-part4/internal"
	"mic-trainning-lessons-part4/model"
	"mic-trainning-lessons-part4/proto/pb"
)

type ShopCartService struct {

}

type CartOrderService struct {

}

func (s ShopCartService) ShopCartItemList(ctx context.Context, req *pb.AccountReq) (*pb.CartItemListRes, error) {
	var cartItemList []model.ShopCart
	var res *pb.CartItemListRes
	var cartItemListPb []*pb.CarItemRes

	r := internal.DB.Where(&model.ShopCart{AccountId: req.AccountId}).Find(&cartItemList)
	if r.Error != nil {
		return nil, errors.New(custom_error.ParamError)
	}
	if r.RowsAffected < 1 {
		return res, nil
	}

	for _, item := range cartItemList {
		itemPb := ConvertShopCarModel2Pb(item)
		cartItemListPb = append(cartItemListPb, itemPb)
	}
	res.ItemList = cartItemListPb
	res.Total = int32(r.RowsAffected)
	return res, nil
}

func ConvertShopCarModel2Pb(s model.ShopCart) *pb.CarItemRes {
	cart := &pb.CarItemRes{
		Id:        s.ID,
		AccountId: s.AccountId,
		ProductId: s.ProductId,
		Num:       int32(s.Num),
		Checked:   s.Checked,
	}
	return cart
}

func (s ShopCartService) AddShopCarItem(ctx context.Context, req *pb.ShopCartReq) (*pb.CarItemRes, error) {

	var cart model.ShopCart

	r := internal.DB.Where(&model.ShopCart{AccountId: req.AccountId, ProductId: req.ProductId}).First(&cart)
	if r.RowsAffected < 1 {
		cart.AccountId = req.AccountId
		cart.ProductId = req.ProductId
		cart.Num = int(req.Num)
		cart.Checked = req.Checked
	} else {
		cart.Num += int(req.Num)
		cart.Checked = req.Checked
	}

	internal.DB.Save(&cart)
	res := ConvertShopCarModel2Pb(cart)
	return res, nil

}

func (s ShopCartService) UpdateShopCartItem(ctx context.Context, req *pb.ShopCartReq) (*emptypb.Empty, error) {
	var cart model.ShopCart

	r := internal.DB.Where(&model.ShopCart{
		AccountId: req.AccountId,
		ProductId: req.ProductId,
	}).Find(&cart)

	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CartNoFount)
	}
	if req.Num < 1 {
		return nil, errors.New(custom_error.ParamError)
	}

	cart.Num = int(req.Num)
	cart.Checked = req.Checked
	internal.DB.Save(&cart)
	return &emptypb.Empty{}, nil
}

func (s ShopCartService) DeleteShopCartItem(ctx context.Context, req *pb.DelShopCartItem) (*emptypb.Empty, error) {

	var cart model.ShopCart

	r := internal.DB.Where(&model.ShopCart{
		AccountId: req.AccountId,
		ProductId: req.ProductId,
	}).Delete(&cart)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CartDelFal)
	}

	return &emptypb.Empty{}, nil
}
