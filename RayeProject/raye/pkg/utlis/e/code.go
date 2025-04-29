package e

const (
	Success               = 0
	Failed                = 1
	InvalidParam          = 2
	Unauthorized          = 3
	NotFound              = 4
	NotLogin              = 401
	OffLine               = 402
	NotLoginSqueeze       = 409
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	DeleteSuccess         = 204
	NotExistInentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400
	InvalidToken          = 501
)

var (
	JWTSecret = []byte("raye1017")
)

const (
	ErrorAuthCheckTokenFail    = 30001 //token 错误
	ErrorAuthCheckTokenTimeout = 30002 //token 过期
)
