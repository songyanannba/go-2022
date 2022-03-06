package hander

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mic-training-lessons-part2/custom_error"
	"mic-training-lessons-part2/internal"
	"mic-training-lessons-part2/proto/pb"
	"net/http"
	"strconv"
)

var productClient pb.ProductServiceClient

func init() {
	addr := fmt.Sprintf("%s:%d" , internal.AppConf.ProductSrvConfig.Host , internal.AppConf.ProductSrvConfig.Port)
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))

	if err != nil {
		panic(err)
	}

	productClient = pb.NewProductServiceClient(conn)

}

func ProductListHandler(c *gin.Context) {
	var condition pb.ProductConditionReq
	minPriceStr := c.DefaultQuery("minPrice", "0")
	minPrice, err := strconv.Atoi(minPriceStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": custom_error.ParamError})
		return
	}
	maxPriceStr := c.DefaultQuery("maxPrice", "0")
	maxPrice, err := strconv.Atoi(maxPriceStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
	}
	condition.MinPrice = int32(minPrice)
	condition.MaxPrice = int32(maxPrice)

	categoryIdStr :=  c.DefaultQuery("categoryId" , "0")
	categoryId ,err := strconv.Atoi(categoryIdStr)
	if err != nil {
		c.JSON(http.StatusOK , gin.H {
			"msg" : custom_error.ParamError,
		})
	}
	condition.CategoryId = int32(categoryId)

	brandStr := c.DefaultQuery("brandId" , "0")
	brandId, err := strconv.Atoi(brandStr)
	if err != nil {
		c.JSON(http.StatusOK , gin.H {
			"msg" : custom_error.ParamError,
		})
	}
	condition.BrandId = int32(brandId)

	isPop := c.DefaultQuery("isPop" , "0")
	if isPop == "1" {
		condition.IsPop = true
	}

	isNew := c.DefaultQuery("isNew" , "0")
	if isNew == "1" {
		condition.IsNew = true
	}

	pageNoStr := c.DefaultQuery("pageNo" , "0")
	pageNo ,err := strconv.Atoi(pageNoStr)
	if err != nil {
		c.JSON(http.StatusOK , gin.H {
			"msg" : custom_error.ParamError,
		})
	}
	condition.PageNo = int32(pageNo)

	pageSizeStr := c.DefaultQuery("pageSize" , "0")
	pageSize ,err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusOK , gin.H {
			"msg" : custom_error.ParamError,
		})
		return
	}
	condition.PageSize = int32(pageSize)

	keyword := c.DefaultQuery("keyword" , "")
	condition.KeyWord = keyword

	list, err := productClient.ProductList(context.Background(), &condition)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK , gin.H {
			"msg" : custom_error.ParamError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "",
		"total" : list.Total,
		"data" : list.ItemList,
	})
}
