package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodePWDNotMatch
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvalidParam: "invalid param",
	CodeUserExist:    "user is already exist",
	CodeUserNotExist: "user does not exist",
	CodePWDNotMatch:  "password not match",
	CodeServerBusy:   "server is busy",

	CodeInvalidToken: "invalid token",
	CodeNeedLogin:    "need login",
}

func (rc ResCode) Msg() string {
	msg, ok := codeMsgMap[rc]
	if !ok {
		msg = "invalid response code"
	}
	return msg
}
