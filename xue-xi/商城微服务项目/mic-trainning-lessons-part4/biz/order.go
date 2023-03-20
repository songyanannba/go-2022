package biz

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part4/custom_error"
	"mic-trainning-lessons-part4/internal"
	"mic-trainning-lessons-part4/model"
	"mic-trainning-lessons-part4/proto/pb"
)

func (c CartOrderService) CreateOrder(ctx context.Context, req *pb.OrderItemReq) (*pb.OrderItemRes, error) {

	//var order model.OrderItem
	//var orderRes *pb.OrderItemRes
	var productIds []int32
	var cartList []model.ShopCart

	productNumMap := make(map[int32]int32)

	//获取购物车里面的商品ID
	r := internal.DB.Where(&model.ShopCart{
		AccountId: req.AccountId, Checked: false,
	}).Find(&cartList)
	if r.RowsAffected == 0 {
		return nil, errors.New(custom_error.CartNotFindProduct)
	}
	for _, cart := range cartList {
		productIds = append(productIds, cart.ProductId)
		productNumMap[cart.ProductId] = int32(cart.Num)
	}

	productRes, err := internal.ProductClient.BatchGetProduct(context.Background(), &pb.BatchProductIdReq{
		Ids: productIds,
	})
	if err != nil {
		return nil, errors.New("ProductClient.BatchGetProduct err")
	}
	var amount float32
	var orderProductList []model.OrderProduct
	var stockItemsList []*pb.ProductStockItem

	for _, p := range productRes.ItemList {
		amount += p.RealPrice * float32(productNumMap[p.Id])
		var orderProduct = model.OrderProduct{
			ProductId:   p.Id,
			ProductName: p.Name,
			CoverImage:  p.CaverImage,
			RealPrice:   p.RealPrice,
			Num:         productNumMap[p.Id],
		}
		orderProductList = append(orderProductList, orderProduct)
		stockItems := &pb.ProductStockItem{
			ProductId: p.Id,
			Num:       productNumMap[p.Id],
		}
		stockItemsList = append(stockItemsList, stockItems)
	}

	_, err = internal.StockClient.Sell(context.Background(), &pb.SellItem{StockItemList: stockItemsList})
	if err != nil {
		return nil, errors.New("internal.StockClient.Sell err")
	}
	tx := internal.DB.Begin()
	orderItem := model.OrderItem{
		AccountId:     req.AccountId,
		OrderNo:       uuid.New().String(),
		Status:        "unPay",
		Addr:          req.Addr,
		Receiver:      req.Receiver,
		ReceiverModel: req.Mobile,
		PostCode:      req.PostCode,
		PayTime:       nil,
	}

	tx.Save(orderItem)

	for _, orderProduct := range orderProductList {
		orderProduct.OrderId = orderItem.ID
	}

	result := tx.CreateInBatches(orderProductList, 50)

	if result.Error != nil || result.RowsAffected < 1 {
		tx.Rollback()
		return nil, errors.New("CreateInBatches err")
	}

	result = tx.Where(&model.ShopCart{Checked: true, AccountId: req.AccountId}).Delete(&model.ShopCart{})
	if result.Error != nil || result.RowsAffected < 1 {
		tx.Rollback()
		return nil, errors.New("C err")
	}
	tx.Commit()

	res := pb.OrderItemRes{
		Id:        orderItem.ID,
		AccountId: orderItem.AccountId,
	}

	return &res, nil
}

func (c CartOrderService) OrderList(ctx context.Context, req *pb.OrderPagingReq) (*pb.OrderListRes, error) {
	var orderList []model.OrderItem
	var res *pb.OrderListRes
	var total int64

	//查询数据库
	//r := internal.DB.Where(&model.OrderItem{AccountId: req.AccountId}).Find(&orderList)
	//res.Total = int32(r.RowsAffected)

	internal.DB.Where(&model.OrderItem{AccountId: req.AccountId}).Count(&total)
	res.Total = int32(total)

	internal.DB.Scopes(internal.MyPaging(int(req.PagNo), int(req.PageSize))).Find(&orderList)

	for _, item := range orderList {
		model2pb := ConvertOrderModel2pb(item)
		res.ItemList = append(res.ItemList, model2pb)
	}

	return res, nil
}

func ConvertOrderModel2pb(o model.OrderItem) *pb.OrderItemRes {
	res := pb.OrderItemRes{
		Id:        o.ID,
		AccountId: o.AccountId,
		PayType:   o.PayType,
		OrderNo:   o.OrderNo,
		PostCode:  o.PostCode,
		Amount:    o.OrderAmount,
		Addr:      o.Addr,
		Receiver:  o.Receiver,
		Mobile:    "",
		Status:    o.Status,
	}
	return &res
}

func (c CartOrderService) OrderDetail(ctx context.Context, req *pb.OrderItemReq) (*pb.OrderItemDetailRes, error) {

	var orderDetail model.OrderItem
	var detailRes *pb.OrderItemDetailRes

	r := internal.DB.Where(&model.OrderItem{BaseModel: model.BaseModel{ID: req.Id}}).First(&orderDetail)

	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.OrderNotFind)
	}

	res := ConvertOrderModel2pb(orderDetail)
	detailRes.Order = res
	var orderProductList []model.OrderProduct
	internal.DB.Where(&model.OrderProduct{OrderId: orderDetail.ID}).First(&orderProductList)

	for _, product := range orderProductList {
		orderProductRes := ConvertOrderProduct2Pb(product)
		detailRes.ProductList = append(detailRes.ProductList, orderProductRes)
	}

	return detailRes, nil
}

func ConvertOrderProduct2Pb(p model.OrderProduct) *pb.OrderProductRes {
	res := &pb.OrderProductRes{
		Id:          p.ID,
		OrderId:     p.OrderId,
		ProductId:   p.ProductId,
		Num:         p.Num,
		ProductName: p.ProductName,
		RealPrice:   p.RealPrice,
		CoverImage:  p.CoverImage,
	}
	return res
}

func (c CartOrderService) ChangOrderStatus(ctx context.Context, req *pb.OrderStatus) (*emptypb.Empty, error) {
	r := internal.DB.Model(&model.OrderItem{}).Where("order_no=?", req.OrderNo).Update("status=?", req.Status)

	if r.RowsAffected == 0 {
		return nil, errors.New(custom_error.ParamError)
	}
	return &emptypb.Empty{}, nil
}
