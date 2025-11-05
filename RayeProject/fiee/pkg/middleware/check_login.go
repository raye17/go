package middleware

import (
	"errors"
	"test001/api/account"
	"test001/api/rule"
	"test001/model"
	"test001/pkg/e"
	"test001/pkg/secret"
	"test001/pkg/service"

	"github.com/gin-gonic/gin"
)

const (
	ErrNotLogin = "请先登录"
)

const (
	Authorization = "Authorization"
)

func CheckLogin(provider *account.AccountClientImpl) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//如果没有登录
		authorization := ctx.GetHeader(Authorization)
		if authorization == "" {
			service.Error(ctx, e.NotLogin, errors.New(ErrNotLogin))
			return
		}

		//jwt, nowDepartmentID, _, err := secret.GetJwtFromStr(authorization)
		jwt, err := secret.GetJwtFromStrV2(authorization)

		if err != nil {
			service.Error(ctx, e.NotLogin, errors.New(ErrNotLogin))
			return
		}

		req := account.DecryptJwtRequest{
			Token: jwt,
		}

		info, err := provider.DecryptJwt(ctx, &req)

		if err != nil {
			service.Error(ctx, e.NotLogin, err)
			return
		}

		//获取用户的账号信息
		infoReq := &account.InfoRequest{
			ID: info.ID,
		}

		infoRes, err := service.AccountProvider.Info(ctx, infoReq)

		if err != nil {
			service.Error(ctx, e.Error, err)
			return
		}

		//获取用户的岗位信息
		uReq := rule.RulesRequest{
			AccountID: info.ID,
		}

		qres, err1 := service.RuleProvider.UserInfo(ctx, &uReq)

		if err1 != nil {
			service.Error(ctx, e.Error, err)
			return
		}

		loginInfo := model.LoginInfo{
			Domain:         info.Domain,
			ID:             info.ID,
			Account:        info.Account,
			NickName:       info.NickName,
			PositionUsers:  qres.PositionUsers,
			Extend:         infoRes.Info.Extend,
			TelNum:         infoRes.Info.TelNum,
			JumpTo:         "",
			DepartmentName: "",
			//NowPositionUser: nil,
		}

		//for _, tt := range qres.PositionUsers {
		//	if tt.DepartmentId == nowDepartmentID {
		//		loginInfo.NowPositionUser = tt
		//		loginInfo.DepartmentName = tt.DepartmentName
		//	}
		//}

		if infoRes.Info.Extend != nil {
			loginInfo.JumpTo = infoRes.Info.Extend.JumpTo
		}

		//if loginInfo.NowPositionUser == nil {
		//	service.Error(ctx, e.NotLogin, errors.New("账号没有绑定部门"))
		//	return
		//}
		//fmt.Println("当前部门是++++", loginInfo.NowPositionUser.DepartmentName)

		ctx.Set("jwtInfo", loginInfo)

		ctx.Next()
	}
}
