// Code generated by goctl. DO NOT EDIT.
package types

type Article struct {
	Title       string `json:"title"`
	ArticleId   int64  `json:"articleId"`
	ImageURL    string `json:"imageURL"`
	Desc        string `json:"desc"`
	PublishTime string `json:"publishTime"`
	UpdateTime  string `json:"updateTime"`
}

type GetArticleListResp struct {
	Data []*Article `json:"data"`
	Meta Meta       `json:"meta"`
}

type GetArticlePathReq struct {
	ParticleId int64 `path:"particleId"`
}

type GetArticlePathResp struct {
	Data string `json:"data"`
	Meta Meta   `json:"meta"`
}

type AddArticleReq struct {
	Title       string `json:"title"`
	MdFileName  int32  `json:"mdFileName"`
	ImageURL    string `json:"imageURL"`
	Desc        string `json:"desc"`
	PublishTime string `json:"publishTime"`
	UpdateTime  string `json:"updateTime"`
}

type AddArticleResp struct {
	ArticleURL string `json:"articleURL"`
}

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
