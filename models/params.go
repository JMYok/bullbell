package models

//定义请求参数的结构体

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required" min:"4" max:"100"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" min:"4" max:"100"`
}

// LoginRes 结果参数结构体
type LoginRes struct {
	Username     string `json:"user_name"`
	UserId       uint64 `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// ParamPostRequest 请求参数结构体
type ParamPostRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	AuthorId    uint64 `json:"author_id,string"`
	CommunityId uint64 `json:"community_id,string" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	//UserID 请求中获取
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=-1 0 1" ` // 赞成票(1) 反对票(-1) 取消投票(0)
}
