package controller

type ResCode int

const (
	CodeSuccess      ResCode = 1000
	CodeInvalidParam ResCode = 1001
	CodeUserExist    ResCode = 1002
	CodeUserNotExist ResCode = 1003
	CodePWDNotMatch  ResCode = 1004
	CodeServerBusy   ResCode = 1005
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:      "success",
	CodeInvalidParam: "invalid param",
	CodeUserExist:    "user is already exist",
	CodeUserNotExist: "user does not exist",
	CodePWDNotMatch:  "password not match",
	CodeServerBusy:   "server is busy",
}

func (rc ResCode) Msg() string {
	msg, ok := codeMsgMap[rc]
	if !ok {
		msg = "invalid response code"
	}
	return msg
}
