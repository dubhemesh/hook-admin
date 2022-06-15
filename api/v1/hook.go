package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateHookReq struct {
	g.Meta `path:"/hook/create" tags:"Hook" method:"post" sm:"创建Hook"`
}

type CreateHookRes struct{}

type DeleteHookReq struct {
	g.Meta `path:"/hook/delete" tags:"Hook" method:"post" sm:"删除Hook"`
	Id     int `json:"id" v:"required" dc:"Hook ID"`
}

type DeleteHookRes struct{}

type UpdateHookReq struct {
}
