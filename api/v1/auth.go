package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Auth" method:"post" sm:"登录"`
	UserName string `json:"user_name" v:"required" dc:"用户名"`
	Password string `json:"password" v:"required" dc:"密码"`
}

type LoginRes struct {
	Token     string `json:"token"`
	ExpiredIn int64  `json:"expired_in"`
}

type RefreshReq struct {
	g.Meta `path:"/refresh_token" tags:"Auth" method:"post" sm:"刷新token"`
	Token  string `json:"token" v:"required" dc:"token" in:"header"`
}

type RefreshRes struct {
	Token     string `json:"token"`
	ExpiredIn int64  `json:"expired_in"`
}

type LogoutReq struct {
	g.Meta `path:"/register" tags:"Auth" method:"post" sm:"注册"`
	Token  string `json:"token" v:"required" dc:"token" in:"header"`
}

type LogoutRes struct{}
