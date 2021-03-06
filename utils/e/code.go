package e

const (
	//通用
	SUCCESS        = 20000
	ERROR          = 50000
	INVALID_PARAMS = 40000

	//用户验证
	ERROR_EXIST_USER     = 10001
	ERROR_NOT_EXIST_USER = 10002
	ERROR_WRONG_PASSWORD = 10003
	ERROR_WRONG_ID       = 10004

	//token验证
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ACCESS_TOKEN                   = 20005
)
