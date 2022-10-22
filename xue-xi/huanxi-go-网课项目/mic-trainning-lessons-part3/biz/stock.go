package biz

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part3/custom_error"
	"mic-trainning-lessons-part3/internal"
	"mic-trainning-lessons-part3/model"
	"mic-trainning-lessons-part3/proro/pb"
	"sync"
)

type StockServer struct {
}

func (s StockServer) SetStock(ctx context.Context, req *pb.ProductStockItem) (*emptypb.Empty, error) {

	//参数校验
	var stock model.Stock
	internal.DB.Where("product_id=?", req.ProductId).Find(&stock)
	if stock.ID < 1 {
		stock.ProductId = req.ProductId
		stock.Num = req.Num
	} else {
		stock.Num = req.Num + stock.Num
	}

	internal.DB.Save(&stock)
	return &emptypb.Empty{}, nil

}

func (s StockServer) StockDetail(ctx context.Context, req *pb.ProductStockItem) (*pb.ProductStockItem, error) {

	var stock model.Stock
	r := internal.DB.Where("product_id=?", req.ProductId).First(stock)

	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.ParamError)
	}

	return ConvertStockModel2Pb(stock), nil
}

func ConvertStockModel2Pb(m model.Stock) *pb.ProductStockItem {
	res := &pb.ProductStockItem{
		ProductId: m.ProductId,
		Num:       m.Num,
	}
	return res
}

var lock sync.Mutex

func (s StockServer) Sell(ctx context.Context, req *pb.SellItem) (*emptypb.Empty, error) {
	//减库存
	tx := internal.DB.Begin()
	lock.Lock() //todo 性能差
	defer lock.Unlock()
	for _, item := range req.StockItemList {
		var stock model.Stock
		r := internal.DB.Where("product_id=?", item.ProductId).First(&stock)
		if r.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New(custom_error.ProductNoFount)
		}
		if stock.Num < item.Num {
			tx.Rollback()
			return nil, errors.New(custom_error.ProductNoFount)
		}
		stock.Num -= item.Num

		tx.Save(stock)
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}

func (s StockServer) BackStock(ctx context.Context, req *pb.SellItem) (*emptypb.Empty, error) {
	/**
	什么时候归还
	1 订单超时
	2 订单创建失败
	3 手动
	 */
	tx := internal.DB.Begin()
	for _, item := range req.StockItemList {
		var stock model.Stock
		r := internal.DB.Where("product_id=?", item.ProductId).First(&stock)
		if r.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New(custom_error.ProductNoFount)
		}
		stock.Num += item.Num
		tx.Save(stock)
	}
	tx.Commit()
	return &emptypb.Empty{}, nil
}
