package e

var (
	JWTSecret = []byte("asdfqwer1234")
)
var (
	ShowStatusNeedSellerSure = 1 //待销售确认收款
	ShowStatusNeedReceive    = 2 //代后台确认收款
	ShowStatusNeedAgreed     = 3 //代协议
	ShowStatusNeedReap       = 4 //代收获
	ShowStatusDone           = 5 //完毕
)

const (
	Success      = 200
	Error        = 500
	ParamsError  = 400
	InvalidToken = 501
)
const (
	AuctionDomain = "auction"
	Pay_Success   = 1
)
const (
	FROM_SELF = "seller"
	FROM_MALL = "tf"
)

const (
	ClientID = "tf"
	SECRET   = "93a8894116ce0152dd5e145370fc83aa"
)

const (
	SellerWeekBossKeyUrl = "r_salebot_index_btn_site" //能看周报的权限
	SellerBossPosition   = "销售总监"
	SellerBossKey        = "seller_mobile_boss_auth"
	SellerDrawKey        = "seller_draw_auth"
	SellerTeamReportKey  = "seller_team_report_auth" //团队报表权限
	SellerReportsAllKey  = "seller_reports_all_auth" //销售宝周报全部站点权限
	SellerStaffKey       = "seller_mobile_staff_auth"
	DefaultAvatar        = "https://dci-file-new.bj.bcebos.com/fonchain-main/test/runtime/image/avatar/40/b8ed6fea-6662-416d-8bb3-1fd8a8197061.jpg"
	TimeDateFormat       = "2006-01-02 15:04:05"
	Setting              = "/setting"
)

const (
	Failed                = 1
	Ok                    = 0
	BindError             = 2
	JsonUnmarshal         = 3
	NotLogin              = 401
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	DeleteSuccess         = 204
	NotExistInentifier    = 202
	ERROR                 = 500
	InvalidParams         = 400

	//成员错误
	ErrorExistNick          = 10001
	ErrorExistUser          = 10002
	ErrorNotExistUser       = 10003
	ErrorNotCompare         = 10004
	ErrorNotComparePassword = 10005
	ErrorFailEncryption     = 10006
	ErrorNotExistProduct    = 10007
	ErrorNotExistAddress    = 10008
	ErrorExistFavorite      = 10009
	ErrorGetUserInfo        = 10010
	ErrorGetDepart          = 10011
	ErrorUpdateAw           = 10012

	//店家错误
	ErrorBossCheckTokenFail        = 20001
	ErrorBossCheckTokenTimeout     = 20002
	ErrorBossToken                 = 20003
	ErrorBoss                      = 20004
	ErrorBossInsufficientAuthority = 20005
	ErrorBossProduct               = 20006

	//管理员错误
	ErrorAuthCheckTokenFail        = 30001 //token 错误
	ErrorAuthCheckTokenTimeout     = 30002 //token 过期
	ErrorAuthToken                 = 30003
	ErrorAuth                      = 30004
	ErrorAuthInsufficientAuthority = 30005
	ErrorReadFile                  = 30006
	ErrorSendEmail                 = 30007
	ErrorCallApi                   = 30008
	ErrorUnmarshalJson             = 30009
	ErrorAdminFindUser             = 30010
	//数据库错误
	ErrorDatabase = 40001

	//对象存储错误
	ErrorOss        = 50001
	ErrorUploadFile = 50002

	//店铺错误
	ErrorExistShopName    = 60001
	ErrorNotExistShopName = 60002
	ErrorNotAdmin         = 60003

	ErrNoDomain      = 70001
	ErrTelNum        = 70002
	ErrNoCode        = 70003
	ErrNoID          = 70004
	ErrNickName      = 70005
	InvalidID        = 70006
	InvalidPas       = 70007
	ErrStatus        = 70008
	ErrNoType        = 70009
	ErrNoUserID      = 70010
	ErrNoName        = 70011
	ErrNoDepCode     = 70012
	ErrNoTitle       = 70013
	ErrNoUrl         = 70014
	ErrNoMethod      = 70015
	ErrNotDep        = 70016
	ErrCreateQr      = 70017
	ErrNotSellerBoss = 70018
	ErrWrongDate     = 70019
	ErrLoginSeller   = 70020

	//上传
	ErrorUploadVideoCover = 80001
	ErrorUploadValidParam = 80002
	ErrorFileReadErr      = 80003
	ErrorFileNotExists    = 80004
	ErrorChunkNotGt       = 80005
	ErrorChunk            = 80006
	ErrorUploadBos        = 80007
	ErrorFileCreate       = 80008
	ERROR_UID             = 80009

	//画作
	ErrorAllotUids = 90001
)
const (
	Push      = 1
	Read      = 2
	NotFilled = 3
	Save      = 4
)
const (
	Audit  = 1
	Pass   = 2
	Reject = 3
)
