package hander

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mic-trainning-lessons/account_srv/proto/pb"
	"mic-trainning-lessons/account_web/req"
	"mic-trainning-lessons/account_web/res"
	"mic-trainning-lessons/custom_error"
	"mic-trainning-lessons/internal"
	"mic-trainning-lessons/jwt_op"
	"mic-trainning-lessons/log"
	"net/http"
	"strconv"
	"time"
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

var (
	client pb.AccountServiceClient
)

func init() {
	err := initGrpcClient()
	if err != nil {
		panic(err)
	}
}

func initGrpcClient() error {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.ConsulConfig.Host, internal.AppConf.ConsulConfig.Port)
	dialAddr := fmt.Sprintf("consul://%s/%s?wait=14", addr, internal.AppConf.AccountSrvConfig.SrvName)
	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Fatal(err)
		s := fmt.Sprintf("AccountListHandler-grpc拨号失败:%s", err.Error())
		log.Logger.Info(s)
		return err
	}
	client = pb.NewAccountServiceClient(conn)
	return nil
}

func AccountListHandler(c *gin.Context) {
	//log.Logger.Info("AccountListHandler调试通过...")

	pageNoStr := c.DefaultQuery("pageNo", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "3")
	pageNo, _ := strconv.ParseInt(pageNoStr, 10, 32)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 32)

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

func LoginByPasswordHandler(c *gin.Context) {
	var loginByPassword req.LoginByPassword
	err := c.ShouldBind(&loginByPassword)
	if err != nil {
		log.Logger.Error("LoginByPassword出错:" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": "解析参数错误",
		})
		return
	}

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

	client := pb.NewAccountServiceClient(conn)

	r, err := client.GetAccountByMobile(context.Background(), &pb.MobileRequest{
		Mobile: loginByPassword.Mobile,
	})
	if err != nil {
		s := fmt.Sprintf("GetAccountByMobile失败:%s", err.Error())
		log.Logger.Info(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"mag": e,
		})
		return
	}

	cheRes, err := client.CheckPassword(context.Background(), &pb.CheckAccountRequest{
		Password:       loginByPassword.Password,
		HashedPassword: r.Password,
		AccountId:      uint32(r.Id),
	})
	if err != nil {
		s := fmt.Sprintf("grpc-CheckPassword失败:%s", err.Error())
		log.Logger.Info(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"mag": e,
		})
		return
	}
	checkResult := "登陆失败"
	if cheRes.Result {
		checkResult = "登陆成功"
		j := jwt_op.NewJWT()
		now := time.Now()
		claims := jwt_op.CustomClaims{
			StandardClaims: jwt.StandardClaims{
				NotBefore: now.Unix(),
				ExpiresAt: now.Add(time.Hour * 24 * 30).Unix(),
			},
			ID:          uint32(r.Id),
			NickName:    r.Nickname,
			AuthorityIs: int32(r.Role),
		}
		token, err := j.GenerateJWT(claims)
		if err != nil {
			s := fmt.Sprintf("GenerateJWT失败:%s", err.Error())
			log.Logger.Info(s)
			e := HandleError(err)
			c.JSON(http.StatusOK, gin.H{
				"mag": e,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"mag":    "",
			"result": checkResult,
			"token":  token,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"mag":    checkResult,
		"result": checkResult,
		"token":  "",
	})
}

//健康检查
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
