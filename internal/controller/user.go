package controller

import (
	"context"

	"hook-admins/api/v1"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	return
}

func (c *cUser) UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	return
}

func (c *cUser) UserRegister(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	return
}

func (c *cUser) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	return
}

func (c *cUser) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	return
}
