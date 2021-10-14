package errmsg

const (
	SUCCSE = 200
	ERROR = 500

	// code = 1000...用户模块错误

)

var codeMsg = map[int]string{
	SUCCSE: "OK",
	ERROR: "FAIL",
}