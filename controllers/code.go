package controllers

type ResCode uint

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeAuthHeaderEmpty
	CodeAuthHeaderWrongFormat
	CodeAuthInvalid
	CodeServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:               "success",
	CodeInvalidParam:          "请求参数错误",
	CodeUserExist:             "用户名已存在",
	CodeUserNotExist:          "用户不存在",
	CodeInvalidPassword:       "用户名或密码错误",
	CodeAuthHeaderEmpty:       "请求头auth字段为空",
	CodeAuthHeaderWrongFormat: "请求头auth格式有误",
	CodeAuthInvalid:           "无效Auth",
	CodeServerBusy:            "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
