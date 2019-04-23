package conf

const (
	AUTH_ERROR = 401
)

var ErrMsg = map[int]string{
	AUTH_ERROR: "token认证失败",
}
