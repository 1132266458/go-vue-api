package errmsg

const (
	SUCCSE = 200
	ERROR = 500

	// code = 100...用户模块错误
	ERROR_USERNAME_USED    = 1001
)

var codeMsg = map[int]string{
	SUCCSE: 				"OK",
	ERROR: 					"FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}