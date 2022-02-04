package hander

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"mic-trainning-lessons/account_srv/proto/pb"
	"mic-trainning-lessons/account_web/res"
	"mic-trainning-lessons/custom_error"
	"mic-trainning-lessons/log"
	"net/http"
	"strconv"
)

func HandleError(err error) string {
	if err != nil {
		switch err.Error() {
		case custom_error.AccountExists:
			return custom_error.AccountExists
		case custom_error.SaltError:
			return custom_error.SaltError
		case custom_error.InternalError:
			return custom_error.InternalError
		case custom_error.AccountNoFount:
			return custom_error.AccountNoFount
		default:
			return custom_error.NoError
		}
	}
	return ""
}

func AccountListHandler(c *gin.Context) {
	//log.Logger.Info("AccountListHandler调试通过...")
	pageNoStr := c.DefaultQuery("pageNo", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "3")

	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		s := fmt.Sprintf("AccountListHandler-grpc拨号失败:%s", err.Error())
		log.Logger.Info(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"mag": e,
		})
		return
	}

	pageNo, _ := strconv.ParseInt(pageNoStr, 10, 32)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 32)

	client := pb.NewAccountServiceClient(conn)
	accountL, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   uint32(pageNo),
		PageSize: uint32(pageSize),
	})
	if err != nil {
		s := fmt.Sprintf("GetAccountList-调用失败：%s", err.Error())
		log.Logger.Info(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"mag": e,
		})
		return
	}

	var resList []res.Account4Res
	for _, item := range accountL.AccountList {
		resList = append(resList, pb2res(item))
	}
	c.JSON(http.StatusOK, gin.H{
		"mag":   "ok",
		"total": accountL.Total,
		"data":  resList,
	})
}

func pb2res(accountRes *pb.AccountRes) res.Account4Res {
	return res.Account4Res{
		Mobile:   accountRes.Mobile,
		NickName: accountRes.Nickname,
		Gender:   accountRes.Gender,
	}
}
