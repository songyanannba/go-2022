package biz

import (
	"context"
	"fmt"
	"mic-trainning-lessons/account_srv/proto/pb"
	"mic-trainning-lessons/internal"
	"testing"
)

func init() {
	internal.InitDB()
}

func TestAccountServer_AddAccount(t *testing.T) {
	accountServer := AccountServer{}

	for i := 0; i < 5; i++ {
		m := fmt.Sprintf("1300000000%d", i)
		res, err := accountServer.AddAccount(context.Background(), &pb.AddAccountRequest{
			Mobile:   m,
			Password: m,
			Nickname: m,
			Gender:   "male",
		})
		if err != err {
			fmt.Println(err)
			continue
		}
		fmt.Println(res.Id)
	}
}

func TestAccountServer_GetAccountList(t *testing.T) {
	accuntServer := AccountServer{}
	res, err := accuntServer.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   1,
		PageSize: 3,
	})
	if err != nil {
		fmt.Println(err)
	}
	for _, account := range res.AccountList {
		fmt.Println(account.Id, account.Mobile)
	}

	fmt.Println("====")

	res, err = accuntServer.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   2,
		PageSize: 3,
	})
	if err != nil {
		fmt.Println(err)
	}
	for _, account := range res.AccountList {
		fmt.Println(account.Id, account.Mobile)
	}
}

func TestAccountServer_GetAccountByMobile(t *testing.T) {
	accountServer := AccountServer{}

	res, err := accountServer.GetAccountByMobile(context.Background(), &pb.MobileRequest{
		Mobile: "13000000007",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Id, res.Mobile)
	}
}

func TestAccountServer_GetAccountById(t *testing.T) {
	accountServer := AccountServer{}
	res, err := accountServer.GetAccountById(context.Background(), &pb.IdRequest{
		Id: 8,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Id, res.Mobile)
	}
}

func TestAccountServer_UpdateAccount(t *testing.T) {
	accountServer := AccountServer{}
	res, err := accountServer.UpdateAccount(context.Background(), &pb.UpdateAccountRequest{
		Id:       1,
		Mobile:   "13000001100",
		Nickname: "hahah",
		Gender:   "female",
		Role:     2,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Result)
}

func TestAccountServer_CheckPassword(t *testing.T) {
	accountServer := AccountServer{}
	password, err := accountServer.CheckPassword(context.Background(), &pb.CheckAccountRequest{
		Password:       "13000000001",
		HashedPassword: "beafe081e7cf31ba49d1219640607977ffe84940d9b5b0bf36291f24fc85afe3",
		AccountId:      2,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password.Result)
}
