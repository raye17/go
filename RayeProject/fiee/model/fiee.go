package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type WorkLog struct {
	Uuid            string `gorm:"column:uuid;type:varchar(50);NOT NULL;primary_key;" json:"id"`
	WorkUuid        string `gorm:"column:work_uuid;type:varchar(50);NOT NULL;index:idx_work_uuid;default:'';comment:作品uuid" json:"workUuid"`
	Title           string `gorm:"column:title;type:varchar(50);NOT NULL;default:'';" json:"title"`
	Content         string `gorm:"column:content;type:varchar(2000);NOT NULL;default:'';" json:"content"`
	WorkCategory    uint8  `gorm:"column:work_category;type:tinyint(1);NOT NULL;default:1;comment: 1 图文 2 视频" json:"workCategory"`
	UpdateTime      string `gorm:"column:update_time;type:varchar(50);NOT NULL;default:'';comment:更新时间" json:"updateTime"`
	WorkStatus      uint8  `gorm:"column:work_status;type:tinyint(1);NOT NULL;default:1;comment: 1 待提交 2 审核中 3 审核失败 4 待艺人确认 5 艺人驳回 6 发布成功 7 发布失败" json:"workStatus"`
	PlatformIDs     string `gorm:"column:platform_ids;type:json;NOT NULL;comment:发布平台ID集合 TIKTOK= 1, YOUTUBE = 2, INS = 3;" json:"platformIDs"`
	ArtistName      string `gorm:"column:artist_name;type:varchar(50);NOT NULL;default:'';comment:艺人名称" json:"artistName"`
	ArtistUuid      string `gorm:"column:artist_uuid;type:varchar(50);NOT NULL;default:'';comment:艺人ID" json:"artistUuid"`
	MediaAccUserIDs string `gorm:"column:media_acc_user_ids;type:json;NOT NULL;comment:自媒体账号user_ids集合" json:"mediaAccUserIDs"`
	MediaNames      string `gorm:"column:media_names;type:varchar(600);NOT NULL;default:'';comment:自媒体账号名称集合" json:"mediaNames"`
	ExtraData       string `gorm:"column:extra_data;type:varchar(600);NOT NULL;default:'';comment:扩展数据" json:"extraData"`
	OperatorName    string `gorm:"column:operator_name;type:varchar(50);NOT NULL;default:'';comment:操作人名称" json:"operatorName"`
	OperatorID      string `gorm:"column:operator_id;type:varchar(50);NOT NULL;default:'';comment:操作人ID" json:"operatorID"`
	ConfirmedAt     int    `gorm:"column:confirmed_at;type:int(11);NOT NULL;default:0;comment:确认时间" json:"confirmedAt"`
	CreatedAt       int    `gorm:"column:created_at;type:int(11);autoCreateTime" json:"createdAt"`
	UpdatedAt       int    `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updatedAt"`
	DeletedAt       soft_delete.DeletedAt
}

func (WorkLog) TableName() string {
	return "cast_work_log"
}

type CostLog struct {
	Uuid          string `gorm:"column:uuid;type:varchar(50);NOT NULL;primary_key;" json:"id"`
	ArtistUuid    string `gorm:"column:artist_uuid;type:varchar(50);NOT NULL;index:idx_artist_uuid;default:'';comment:艺人ID" json:"artistUuid"`
	ArtistName    string `gorm:"column:artist_name;type:varchar(50);NOT NULL;default:'';comment:艺人名称" json:"artistName"`
	ArtistPhone   string `gorm:"column:artist_phone;type:varchar(50);NOT NULL;default:'';comment:艺人手机号" json:"artistPhone"`
	WorkUuid      string `gorm:"column:work_uuid;type:varchar(50);NOT NULL;index:idx_work_uuid;default:'';comment:作品uuid" json:"workUuid"`
	WorkCategory  uint8  `gorm:"column:work_category;type:tinyint(1);NOT NULL;default:1;comment: 1 图文 2 视频" json:"workCategory"`
	BundleUuid    string `gorm:"column:bundle_uuid;type:varchar(50);NOT NULL;index:idx_bundle_uuid;default:'';comment:套餐ID uuid" json:"bundleUuid"`
	BundleName    string `gorm:"column:bundle_name;type:varchar(50);NOT NULL;default:'';comment:套餐名称" json:"bundleName"`
	PlatformIds   string `gorm:"column:platform_ids;type:json;NOT NULL;comment:发布平台ID集合" json:"platformIDs"`
	MediaNames    string `gorm:"column:media_names;type:varchar(600);NOT NULL;default:'';comment:自媒体账号名称集合" json:"mediaNames"`
	MediaAccIDs   string `gorm:"column:media_acc_ids;type:varchar(600);NOT NULL;default:'';comment:自媒体账号ID集合" json:"mediaAccIDs"`
	WorkTitle     string `gorm:"column:work_title;type:varchar(50);NOT NULL;" json:"workTitle"`
	SubmitTime    string `gorm:"column:submit_time;type:varchar(50);NOT NULL;default:'';comment:提交时间" json:"submitTime"`
	OperatorName  string `gorm:"column:operator_name;type:varchar(50);NOT NULL;default:'';comment:操作人名称" json:"operatorName"`
	OperatorID    string `gorm:"column:operator_id;type:varchar(50);NOT NULL;default:'';comment:操作人ID" json:"operatorID"`
	OperatorPhone string `gorm:"column:operator_phone;type:varchar(50);NOT NULL;default:'';comment:操作人手机号" json:"operatorPhone"`
	Status        uint8  `gorm:"column:status;type:tinyint(1);NOT NULL;default:1;comment: 1 有效 2 失效" json:"status"`
	CreatedAt     int    `gorm:"column:created_at;type:int(11);autoCreateTime" json:"createdAt"`
	UpdatedAt     int    `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updatedAt"`
	DeletedAt     soft_delete.DeletedAt
}

func (CostLog) TableName() string {
	return "cast_cost_log"
}

type User struct {
	ID               uint `gorm:"primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11)" json:"deletedAt"`
	Domain           *string               `gorm:"size:50"`
	SubNum           string                `gorm:"column:sub_num;comment:用户编号" json:"subNum"`
	TelNum           string                `gorm:"column:tel_num;comment:" json:"telNum"`
	TelAreaCode      string                `gorm:"column:tel_area_code;comment:手机区号" json:"telAreaCode"`
	Status           int                   `gorm:"column:status;comment:状态 1:未实名 2:审核中 3:审核失败 4:审核通过" json:"status"`
	RegistrationTime string                `gorm:"column:registration_time;comment:注册时间" json:"registrationTime"`
	AuditTime        string                `gorm:"column:audit_time;comment:审核时间" json:"auditTime"`
	RealNameID       *uint
	RealName         *RealName `gorm:"foreignKey:RealNameID" json:"RealName"`
	PasswordDigest   string
	NotPassRemarks   string `gorm:"column:not_pass_remarks;comment:不通过备注" json:"notPassRemarks"`
	Nickname         string `gorm:"column:nickname;comment:昵称" json:"nickname"`
	Language         string `gorm:"column:language;comment:语言" json:"language"`
	SubscriberNumber string `gorm:"column:subscriber_number;comment:用户id" json:"subscriberNumber"`
}

func (User) TableName() string {
	return "user"
}

// RealName 实名认证模型
type RealName struct {
	ID                 uint `gorm:"primarykey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11)"`
	Name               string                `gorm:"column:name;comment:姓名" json:"name"`
	Sex                string                `gorm:"column:sex;comment:" json:"sex"`
	Nationality        string                `gorm:"column:nationality;comment:国籍" json:"nationality"`
	DocumentType       int                   `gorm:"column:document_type;comment:证件类型:1护照 2身份证 3驾驶证 4居住证 5自拍照 6社保卡" json:"documentType"`
	CertificatePicture string                `gorm:"type:varchar(500);column:certificate_picture;comment:证件照片" json:"certificatePicture"`
	Validity           string                `gorm:"column:validity;comment:证件有效期" json:"validity"`
	PlaceOfResidence   string                `gorm:"column:place_of_residence;comment:居住地" json:"placeOfResidence"`
	GroupPhoto         string                `gorm:"column:group_photo;comment:证件合影" json:"groupPhoto"`
	Attachment         string                `gorm:"column:attachment;comment:附件" json:"attachment"`
}

const (
	WorkStatusPending          = 1 // 待提交
	WorkStatusReviewing        = 2 // 审核中
	WorkStatusReviewFailed     = 3 // 审核失败
	WorkStatusArtistConfirming = 4 // 待艺人确认
	WorkStatusArtistRejected   = 5 // 艺人驳回
	WorkStatusPublishSuccess   = 6 // 发布成功
	WorkStatusPublishFailed    = 7 // 发布失败
)

type PublishConfig struct {
	PublicType uint32 `json:"publicType"` // 公开范围1 公开 2 不必开 3私享
	CanJoin    uint32 `json:"canJoin"`    // 是否允许合拍1 允许 2 不允许
	CanQuote   uint32 `json:"canQuote"`   // 是否允许引用 1允许2 不允许
	CanComment uint32 `json:"canComment"` // 是否允许评论1.允许2 不允许
	IsAI       uint32 `json:"isAI"`       // 是否AI 1 允许 2 不允许
}

type Work struct {
	Uuid                string `gorm:"column:uuid;type:varchar(50);NOT NULL;primary_key;" json:"id"`
	ArtistUuid          string `gorm:"column:artist_uuid;type:varchar(50);NOT NULL;default:'';comment:艺人ID" json:"artistUuid"`
	MediaAccUserIDs     string `gorm:"column:media_acc_user_ids;type:json;NOT NULL;comment:自媒体账号user_ids集合" json:"mediaAccUserIDs"`
	MediaNames          string `gorm:"column:media_names;type:varchar(600);NOT NULL;default:'';comment:自媒体账号名称集合" json:"mediaNames"`
	ArtistName          string `gorm:"column:artist_name;type:varchar(50);NOT NULL;default:'';comment:艺人名称" json:"artistName"`
	ArtistPhone         string `gorm:"column:artist_phone;type:varchar(50);NOT NULL;default:'';comment:艺人手机号" json:"artistPhone"`
	ArtistPhoneAreaCode string `gorm:"column:artist_phone_area_code;type:varchar(50);NOT NULL;default:'';comment:'艺人手机号区号'" json:"artistPhoneAreaCode"`
	WorkCategory        uint8  `gorm:"column:work_category;type:tinyint(1);NOT NULL;default:1;comment: 1 图文 2 视频" json:"workCategory"`
	SubmitTime          string `gorm:"column:submit_time;type:varchar(50);NOT NULL;default:'';comment:提交时间" json:"submitTime"`
	StatusUpdateTime    string `gorm:"column:status_update_time;type:varchar(50);NOT NULL;default:'';comment:状态更新时间" json:"statusUpdateTime"`
	PlatformIDs         string `gorm:"column:platform_ids;type:json;NOT NULL;comment:发布平台ID集合 TIKTOK= 1, YOUTUBE = 2, INS = 3;" json:"platformIDs"`
	Status              uint8  `gorm:"column:status;type:tinyint(1);NOT NULL;default:1;comment: 1 待提交 2 审核中 3 审核失败 4 待艺人确认 5 艺人驳回 6 发布成功 7 发布失败" json:"status"`
	Title               string `gorm:"column:title;type:varchar(50);NOT NULL;" json:"title"`
	Content             string `gorm:"column:content;type:varchar(2000);NOT NULL;" json:"content"`
	Cost                uint8  `gorm:"column:cost;type:tinyint(1);NOT NULL;default:0;comment: 1 消耗 2 未消耗" json:"cost"`
	ConfirmRemark       string `gorm:"column:confirm_remark;type:varchar(300);NOT NULL;default:'';comment:艺人确认备注" json:"confirmRemark"`
	PublishConfig       string `gorm:"column:publish_config;type:varchar(600);NOT NULL;comment:发布配置" json:"publishConfig"`
	ApprovalID          string `gorm:"column:approval_id;type:varchar(50);NOT NULL;default:'';comment:审批ID" json:"approvalID"`
	Source              uint8  `gorm:"column:source;type:tinyint(1);NOT NULL;default:1;comment: 1 系统 2 导入" json:"source"`
	CreatedAt           int    `gorm:"column:created_at;type:int(11);autoCreateTime" json:"createdAt"`
	UpdatedAt           int    `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updatedAt"`
	DeletedAt           soft_delete.DeletedAt
	WorkImage           *WorkImage `gorm:"foreignKey:WorkUuid;references:Uuid" json:"workImage"`
	WorkVideo           *WorkVideo `gorm:"foreignKey:WorkUuid;references:Uuid" json:"workVideo"`
}

func (Work) TableName() string {
	return "cast_work"
}

type WorkImage struct {
	Uuid      string `gorm:"column:uuid;type:varchar(50);NOT NULL;primary_key;" json:"id"`
	WorkUuid  string `gorm:"column:work_uuid;type:varchar(50);NOT NULL;index:idx_work_uuid;default:'';comment:作品uuid" json:"workUuid"`
	ImageUrls string `gorm:"column:image_urls;type:varchar(3000);NOT NULL;default:'';comment:图片url集合" json:"imageUrls"`
	CreatedAt int    `gorm:"column:created_at;type:int(11);autoCreateTime" json:"createdAt"`
	UpdatedAt int    `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt
}

func (WorkImage) TableName() string {
	return "cast_work_image"
}

type WorkVideo struct {
	Uuid      string `gorm:"column:uuid;type:varchar(50);NOT NULL;primary_key;" json:"id"`
	WorkUuid  string `gorm:"column:work_uuid;type:varchar(50);NOT NULL;index:idx_work_uuid;default:'';comment:作品uuid" json:"workUuid"`
	CoverUrl  string `gorm:"column:cover_url;type:varchar(1500);NOT NULL;default:'';comment:封面url" json:"coverUrl"`
	VideoUrl  string `gorm:"column:video_url;type:varchar(1500);NOT NULL;default:'';comment:视频url" json:"videoUrl"`
	CreatedAt int    `gorm:"column:created_at;type:int(11);autoCreateTime" json:"createdAt"`
	UpdatedAt int    `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt
}

func (WorkVideo) TableName() string {
	return "cast_work_video"
}

type WorkExtra struct {
	WorkUuid            string `gorm:"column:work_uuid;type:varchar(50);NOT NULL;index:idx_work_uuid;default:'';comment:作品uuid;primary_key" json:"workUuid"`
	ArtistConfirmedTime int64  `gorm:"column:artist_confirmed_time;type:bigint(20);NOT NULL;default:0;comment:艺人确认时间" json:"artistConfirmedTime"`
	CreatedAt           int    `gorm:"column:created_at;type:int(11);autoCreateTime" json:"createdAt"`
	UpdatedAt           int    `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updatedAt"`
	DeletedAt           soft_delete.DeletedAt
}

func (WorkExtra) TableName() string {
	return "cast_work_extra"
}

type BundleOrderRecords struct {
	gorm.Model
	UUID                  string                `json:"uuid" gorm:"column:uuid;type:varchar(1024);comment:UUID"`
	OrderNo               string                `json:"orderNo" gorm:"column:order_no;type:varchar(1024);comment:交易编号"`
	BundleUUID            string                `json:"bundleUUID" gorm:"column:bundle_uuid;type:varchar(1024);comment:套餐UUID"`
	BundleName            string                `json:"bundleName" gorm:"column:bundle_name;type:varchar(2048);comment:套餐名"`
	CustomerID            string                `json:"customerID" gorm:"column:customer_id;type:varchar(1024);comment:客户ID"`
	CustomerNum           string                `json:"customerNum" gorm:"column:customer_num;type:varchar(1024);comment:客户编号"`
	CustomerName          string                `json:"customerName" gorm:"column:customer_name;type:varchar(1024);comment:客户名"`
	Amount                float32               `json:"amount" gorm:"column:amount;type:decimal(12,2);comment:套餐金额"`
	AmountType            int64                 `json:"amountType" gorm:"column:amount_type;type:int;comment:金额类型"`
	ValueAddBundleUUID    string                `json:"valueAddBundleUUID" gorm:"column:value_add_bundle_uuid;type:varchar(1024);comment:增值套餐记录UUID"`
	ValueAddBundleAmount  float32               `json:"valueAddBundleAmount" gorm:"column:value_add_bundle_amount;type:decimal(12,2);comment:增值套餐金额"`
	ValueAddOriginalPrice float32               `json:"valueAddOriginalPrice" gorm:"column:value_add_original_price;type:decimal(12,2);comment:原单价"`
	ValueAddDiscountPrice float32               `json:"valueAddDiscountPrice" gorm:"column:value_add_discount_price;type:decimal(12,2);comment:优惠单价"`
	ValueAddSavedAmount   float32               `json:"valueAddSavedAmount" gorm:"column:value_add_saved_amount;type:decimal(12,2);comment:节省金额"`
	TotalAmount           float32               `json:"totalAmount" gorm:"column:total_amount;type:decimal(12,2);comment:总金额"`
	Num                   int32                 `json:"num" gorm:"column:num;type:int;comment:视频数量"`
	SignContract          string                `json:"signContract" gorm:"column:sign_contract;type:varchar(1024);comment:签约合同"`
	Signature             string                `json:"signature" gorm:"column:signature;type:text;comment:签字"`
	SignedTime            string                `json:"signedTime" gorm:"column:signed_time;type:varchar(1024);comment:签约时间(北京时间)"`
	PayType               int64                 `json:"payType" gorm:"column:pay_type;type:int;comment:支付类型"`
	PayTime               string                `json:"payTime" gorm:"column:pay_time;type:varchar(1024);comment:支付时间(北京时间)"`
	CheckoutSessionId     string                `json:"checkoutSessionId" gorm:"column:checkout_session_id;type:varchar(1024);default:null;comment:checkoutSessionId"`
	CheckoutSessionUrl    string                `json:"checkoutSessionUrl" gorm:"column:checkout_session_url;type:varchar(1024);default:null;comment:checkoutSessionUrl"`
	Status                int64                 `json:"status" gorm:"column:status;type:int;comment:状态 1:已签未支付 2:已签已支付"`
	ContractNo            string                `json:"contractNo" gorm:"column:contract_no;type:varchar(1024);comment:合同编号"`
	BundleCommonUid       string                `json:"bundleCommonUid" gorm:"column:bundle_common_uid;type:text;comment:套餐公共ID"`
	AddBundleCommonUid    string                `json:"addBundleCommonUid" gorm:"column:add_bundle_common_uid;type:text;comment:附加套餐公共ID"`
	FinancialConfirmation int32                 `json:"financialConfirmation" gorm:"column:financial_confirmation;type:int;comment:财务确认 1:未确认 2:已确认"`
	ExpirationTime        string                `json:"expirationTime" gorm:"column:expiration_time;comment:套餐过期时间"`
	BundleCommonJson      json.RawMessage       `json:"bundle_common_json" gorm:"column:bundle_common_json;type:json;serializer:json;comment:套餐信息"`
	Language              string                `gorm:"column:language;comment:语言" json:"language"`
	BundleOrderValueAdd   []BundleOrderValueAdd `gorm:"foreignKey:OrderUUID;references:UUID" json:"bundleOrderValueAdd"`
}
type BundleOrderValueAdd struct {
	gorm.Model
	UUID               string  `json:"uuid" gorm:"column:uuid;type:varchar(1024);comment:UUID"`
	OrderNo            string  `json:"orderNo" gorm:"column:order_no;type:varchar(1024);comment:交易编号"`
	OrderUUID          string  `json:"orderUUID" gorm:"column:order_uuid;type:varchar(1024);comment:套餐UUID"`
	CustomerID         string  `json:"customerID" gorm:"column:customer_id;type:varchar(1024);comment:客户ID"`
	CustomerNum        string  `json:"customerNum" gorm:"column:customer_num;type:varchar(1024);comment:客户编号"`
	CustomerName       string  `json:"customerName" gorm:"column:customer_name;type:varchar(1024);comment:客户名"`
	ServiceType        int32   `json:"serviceType" gorm:"column:service_type;type:int;comment:服务类型 1:视频 2:图文 3:数据报表 4:账号数 5:可用时长"`
	CurrencyType       int64   `json:"currencyType" gorm:"column:currency_type;type:int;comment:货币类型"`
	Amount             float64 `json:"amount" gorm:"column:amount;type:decimal(12,2);comment:金额"`
	Num                int32   `json:"num" gorm:"column:num;type:int;comment:数量"`
	Unit               string  `json:"unit" gorm:"column:unit;type:varchar(1024);comment:单位 1个 2条 3天 4月 5年"`
	ValueAddUUID       string  `json:"valueAddUUID" gorm:"column:value_add_uuid;type:varchar(1024);comment:增值服务UUID"`
	Source             int     `json:"source" gorm:"column:source;comment:增加方式 1套餐 2单独购买 3拓展"`
	Remark             string  `json:"remark" gorm:"column:remark;comment:备注"`
	PaymentStatus      int     `json:"paymentStatus" gorm:"column:payment_status;comment:支付状态 1未支付 2已支付"`
	PaymentTime        string  `gorm:"column:payment_time;comment:支付时间" json:"paymentTime"`
	SignContract       string  `json:"signContract" gorm:"column:sign_contract;type:varchar(1024);comment:签约合同"`
	Signature          string  `json:"signature" gorm:"column:signature;type:text;comment:签字"`
	SignedTime         string  `json:"signedTime" gorm:"column:signed_time;type:varchar(1024);comment:签约时间(北京时间)"`
	Snapshot           string  `gorm:"column:snapshot;comment:快照" json:"snapshot"` //订单快照
	CheckoutSessionId  string  `json:"checkoutSessionId" gorm:"column:checkout_session_id;type:varchar(1024);default:null;comment:checkoutSessionId"`
	CheckoutSessionUrl string  `json:"checkoutSessionUrl" gorm:"column:checkout_session_url;type:varchar(1024);default:null;comment:checkoutSessionUrl"`
	HandlingFee        string  `gorm:"column:handling_fee;comment:手续费" json:"handlingFee"`
}
