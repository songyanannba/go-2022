package hander

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"mic-trainning-lessons-part4/cartorder_web/req"
	"mic-trainning-lessons-part4/custom_error"
	"mic-trainning-lessons-part4/internal"
	"mic-trainning-lessons-part4/proto/pb"
	"mic-trainning-lessons-part4/util/otgrpc"
	"net/http"
	"strconv"
)

var shopCartClient pb.ShopCarServiceClient

func init() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.CartOrderSrvConfig.Host, internal.AppConf.CartOrderSrvConfig.Port)
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)

	if err != nil {
		panic(err)
	}

	shopCartClient = pb.NewShopCarServiceClient(conn)

}

func ShopCartListHandler(c *gin.Context) {
	//var condition pb.ShopCartReq
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
		//"total": list.Total,
		//"data":  list.ItemList,
	})
}

func AddHandler(c *gin.Context) {
	var shopCartReq req.ShopCartReq
	err := c.ShouldBindJSON(shopCartReq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
		return
	}

	req2Pb := ConvertProductReq2Pb(shopCartReq)
	shopCart, err := shopCartClient.AddShopCarItem(context.Background(), req2Pb)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mag":  "",
		"data": shopCart,
	})
}

func ConvertProductReq2Pb(shopCartReq req.ShopCartReq) *pb.ShopCartReq {
	item := pb.ShopCartReq{
		//Id:        shopCartReq.,
		//AccountId: shopCartReq,
		//ProductId: 0,
		//Num:       0,
		//Checked:   false,
	}
	return &item
}

func DeleteHandler(c *gin.Context) {
	AccountId := c.Param("account_id")
	accountId, err := strconv.Atoi(AccountId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"mag": custom_error.ParamError,
		})
	}

	ProductId := c.Param("product_id")
	productId, err := strconv.Atoi(ProductId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"mag": custom_error.ParamError,
		})
	}

	shopCart, err := shopCartClient.DeleteShopCartItem(context.Background(),
		&pb.DelShopCartItem{
			AccountId: int32(accountId),
			ProductId: int32(productId),
		})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"mag": custom_error.ParamError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mag":  "",
		"data": shopCart,
	})
}

func UpdateHandler(c *gin.Context) {

	var shopCartReq req.ShopCartReq
	err := c.ShouldBindJSON(shopCartReq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
		return
	}

	req2Pb := ConvertProductReq2Pb(shopCartReq)
	product, err := shopCartClient.UpdateShopCartItem(context.Background(), req2Pb)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mag":  "",
		"data": product,
	})

}
