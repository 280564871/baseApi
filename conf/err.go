package conf

const (
	SUCCESS    = 200
	AUTH_ERROR = 401
	ERROR      = 400

	WX_ERROR = 4001
)

var ErrMsg = map[int]string{
	SUCCESS:    "success",
	AUTH_ERROR: "token认证失败",
	ERROR:      "fail",
	WX_ERROR:   "微信授权出错",
}

func GetMsg(code int) string {
	msg, ok := ErrMsg[code]
	if ok {
		return msg
	}

	return ErrMsg[ERROR]
}
