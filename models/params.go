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
