package hander

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mic-trainning-lessons-part4/internal"
	"mic-trainning-lessons-part4/jwt_op"
	"mic-trainning-lessons-part4/proto/pb"
	"net/http"
	"strconv"
)

func OrderListHandler (c *gin.Context) {
	//订单列表
	accountId , _:= c.Get("accountId")
	claims , _ := c.Get("claims")

	var pageNoStr = c.DefaultQuery("pageNo", "0")
	pageNo , _ := strconv.Atoi(pageNoStr)

	var pageSizeStr = c.DefaultQuery("pageSize", "0")
	pageSize  ,_ := strconv.Atoi(pageSizeStr)

	reqPb := &pb.OrderPagingReq{
		PagNo:     int32(pageNo),
		PageSize:  int32(pageSize),
	}

	customClaims := claims.(*jwt_op.CustomClaims)
	if customClaims.AuthorityIds == 1 {
	 reqPb.AccountId = int32(accountId.(uint))
	}

	ctx := context.WithValue(context.Background() , "webContext" , c)

	r, err := internal.OrderClient.OrderList(ctx, reqPb)
	if err != nil {
		zap.S().Errorw("订单列表查询失败")
		c.JSON(http.StatusOK , gin.H{
			"msg" : "订单列表查询失败",
		})
		return
	}


	c.JSON(http.StatusOK , gin.H{
		"msg" :   "",
		"total" : r.Total,
		"list" : r.ItemList,
	})

}


func AddOrder (c *gin.Context) {
	//orderReq :=

	//internal.OrderClient.CreateOrder()

}

func DetailOrder (c *gin.Context) {


	//internal.OrderClient.OrderDetail()
}