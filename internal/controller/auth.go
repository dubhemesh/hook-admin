package controller

import (
	"context"

	"hook-admins/api/v1"
)

var (
	Hello = cAuth{}
)

type cAuth struct{}

func (c *cAuth) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return
}

func (c *cAuth) Refresh(ctx context.Context, req *v1.RefreshReq) (res *v1.RefreshRes, err error) {
	return
}

func (c *cAuth) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return
}
