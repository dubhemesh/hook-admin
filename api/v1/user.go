package v1

import "github.com/gogf/gf/v2/frame/g"

type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"用户" method:"post" sm:"登录用户"`
	Passport string `json:"passport" v:"required" dc:"用户名"`
	Password string `json:"password" v:"required" dc:"密码"`
}

type UserLoginRes struct{}

type UserLogoutReq struct {
	g.Meta `path:"/user/logout" tags:"用户" method:"post" sm:"注销用户"`
}

type UserLogoutRes struct{}

type UserRegisterReq struct {
	g.Meta          `path:"/user/register" tags:"用户" method:"post" sm:"注册用户"`
	Passport        string `json:"passport" v:"required" dc:"用户名"`
	NickName        string `json:"nick_name" v:"required" dc:"昵称"`
	Password        string `json:"password" v:"required" dc:"密码"`
	ConfirmPassword string `json:"confirm_password" v:"required|same:Password" dc:"确认密码"`
}

type UserRegisterRes struct{}

type UserIsLoginReq struct {
	g.Meta `path:"/user/is_login" tags:"用户" method:"post" sm:"判断用户是否登录"`
}

type UserIsLoginRes struct {
	IsLogin bool `json:"is_login"`
}

type UserIsAdminReq struct {
	g.Meta `path:"/user/is_admin" tags:"用户" method:"post" sm:"用户是否是管理员"`
}

type UserIsAdminRes struct {
	IsAdmin bool `json:"is_admin"`
}

type UpdateUserReq struct {
	g.Meta          `path:"/user/update" tags:"用户" method:"post" sm:"更新用户信息"`
	UserId          int    `json:"user_id" v:"required" dc:"用户ID"`
	NickName        string `json:"nick_name" v:"required" dc:"昵称"`
	Password        string `json:"password" v:"required" dc:"密码"`
	ConfirmPassword string `json:"confirm_password" v:"required|same:Password" dc:"确认密码"`
}

type UpdateUserRes struct{}

type GetUserReq struct {
	g.Meta `path:"/user" tags:"用户" method:"get" sm:"获取用户信息"`
	UserId int `json:"user_id" v:"required" dc:"用户ID"`
}

type GetUserRes struct {
	Id       int    `json:"id" dc:"用户ID"`
	Passport string `json:"passport" dc:"用户名"`
	NickName string `json:"nick_name" dc:"昵称"`
}

type GetUserListReq struct {
	g.Meta `path:"/user/list" tags:"用户" method:"get" sm:"获取用户列表" dc:"管理员权限接口"`
}

type GetUserListRes struct {
	List []*GetUserRes `json:"list" dc:"用户列表"`
}
