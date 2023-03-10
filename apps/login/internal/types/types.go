// Code generated by goctl. DO NOT EDIT.
package types

type Meta struct {
	Status int32  `json:"status,default=200"`
	Msg    string `json:"msg,default=succeed"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespData struct {
	Token string `json:"token"`
}

type LoginResp struct {
	Data LoginRespData `json:"data"`
	Meta Meta          `json:"meta"`
}
