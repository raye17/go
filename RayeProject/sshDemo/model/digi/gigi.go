package digi

import "gorm.io/plugin/soft_delete"

type DciUser struct {
	Id         int64  `gorm:"column:id;type:bigint(20);NOT NULL;primary_key;" json:"id"`
	ArtistUuid string `gorm:"column:artist_uuid;type:varchar(40);comment:画家uuid;NOT NULL;default:'';" json:"artistUuid"`

	CertName          string `gorm:"column:cert_name;type:varchar(40);comment:证件名称;NOT NULL;default:'';" json:"certName"`
	CertificateNumber string `gorm:"column:certificate_number;type:varchar(40);comment:证件号码;NOT NULL;default:'';" json:"certificateNumber"`
	CertificateType   string `gorm:"column:certificate_type;type:varchar(40);comment:证件类型;NOT NULL;default:'';" json:"certificateType"`
	//CertificateFrontFileId string `gorm:"column:certificate_front_file_id;type:varchar(300);comment:证件正面文件路径;NOT NULL;default:'';" json:"certificateFrontFileId"`
	//CertificateBackFileId  string `gorm:"column:certificate_back_file_id;type:varchar(300);comment:证件反面文件路径;NOT NULL;default:'';" json:"certificateBackFileId"`
	CardFaceUrl         string `gorm:"column:card_face_url;type:varchar(300);comment:人脸;NOT NULL;default:'';" json:"cardFaceUrl"`
	CardNationUrl       string `gorm:"column:card_nation_url;type:varchar(300);comment:国徽;NOT NULL;default:'';" json:"cardNationUrl"`
	LegalPersonCertName string `gorm:"column:legal_person_cert_name;type:varchar(40);comment:法人名称;NOT NULL;default:'';" json:"legalPersonCertName"`
	LegalPersonCertType string `gorm:"column:legal_person_cert_type;type:varchar(40);comment:法人证件类型;NOT NULL;default:'';" json:"legalPersonCertType"`
	LegalPersonCertNo   string `gorm:"column:legal_person_cert_no;type:varchar(40);comment:法人证件号;NOT NULL;default:'';" json:"legalPersonCertNo"`
	Phone               string `gorm:"column:phone;type:varchar(40);comment:手机号;NOT NULL;default:'';" json:"phone"`
	AreaType            string `gorm:"column:area_type;type:varchar(40);comment:所属地区;NOT NULL;default:'';" json:"areaType"`

	DciUserId     string                `gorm:"column:dci_user_id;type:varchar(40);comment:dci用户ID;NOT NULL;default:'';" json:"dciUserId"`
	DciUserStatus string                `gorm:"column:dci_user_status;type:varchar(20);comment:dci用户状态;NOT NULL;default:'';" json:"dciUserStatus"`
	Remark        string                `gorm:"column:remark;type:varchar(1000);comme    nt:remark;NOT NULL;default:'';" json:"remark"`
	CreatedAt     int                   `gorm:"column:created_at;type:int(11);autoCreateTime" json:"created_at"`
	UpdatedAt     int                   `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `gorm:"column:deleted_at" json:"DeletedAt"`
}

func (DciUser) TableName() string {
	return "dci_user"
}

type DigitalInfo struct {
	Id                            string `gorm:"column:id;type:varchar(50);NOT NULL;primary_key;" json:"id"`
	PayStatus                     string `gorm:"column:pay_status;type:varchar(20);comment:支付状态;not null;default:'';" json:"payStatus"`
	DigitalRegisterApplyTime      int    `gorm:"column:digital_register_apply_time;type:int(11);comment:数登申请时间;not null;default:0;" json:"digitalRegisterApplyTime"`
	ApplyFormUrl                  string `gorm:"column:apply_form_url;type:varchar(300);comment:用户申请表下载链接;not null;default:'';" json:"applyFormUrl"`
	FlowNumber                    string `gorm:"column:flow_number;type:varchar(100);comment:数登流水号;not null;default:'';" json:"flowNumber"`
	DigitalRegisterCertUrl        string `gorm:"column:digital_register_cert_url;type:varchar(300);comment:数登证书;not null;default:'';" json:"digitalRegisterCertUrl"`
	RegNumber                     string `gorm:"column:reg_number;type:varchar(100);comment:数登登记号;not null;default:'';" json:"regNumber"`
	DigitalRegisterCompletionTime int    `gorm:"column:digital_register_completion_time;type:bigint(20);comment:数登完成时间;not null;default:0;" json:"digitalRegisterCompletionTime"`
	InvoiceUrlList                string `gorm:"column:invoice_url_list;type:varchar(1000);comment:发票;not null;default:'';" json:"invoiceUrlList"`
	DigitalWorkSampleUrl          string `gorm:"column:digital_work_sample_url;type:varchar(300);comment:数登证书;not null;default:'';" json:"digitalWorkSampleUrl"`
	CopyrightId                   string `gorm:"column:copyright_id;type:varchar(50);comment:版权ID;NOT NULL;index:idx_index_copyright_id;default:'';" json:"copyrightId"`
	//Amount                        float32               `gorm:"column:amount;type:decimal(10,4);comment:金额;NOT NULL;default:0;" json:"amount"`
	CreatedAt int                   `gorm:"column:created_at;type:int(11);autoCreateTime" json:"created_at"`
	UpdatedAt int                   `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at" json:"DeletedAt"`
}

func (DigitalInfo) TableName() string {
	return "digital_info"
}

type ArtworkCopyright struct {
	Id          string `gorm:"column:id;type:varchar(50);NOT NULL;primary_key;" json:"id"`
	ArtistUuid  string `gorm:"column:artist_uuid;type:varchar(40);comment:画家uuid;NOT NULL;default:'';" json:"artistUuid"`
	ArtworkUuid string `gorm:"column:artwork_uuid;type:varchar(40);comment:画作uuid;NOT NULL;default:'';index:idx_aw_uuid" json:"artworkUuid"`
	Tfnum       string `gorm:"column:tfnum;type:varchar(40);comment:画作编号;NOT NULL;default:'';index:idx_tfnum" json:"tfnum"`
	WorkName    string `gorm:"column:work_name;type:varchar(100);comment:作品名称;NOT NULL;default:'';" json:"workName"`
	DciUserId   string `gorm:"column:dci_userId;type:varchar(40);comment:dci用户ID;NOT NULL;default:'';" json:"dciUserId"`

	AuthorName      string `gorm:"column:author_name;type:varchar(100);comment:作者姓名;NOT NULL;default:'';" json:"authorName"`
	AuthorSignature string `gorm:"column:author_signature;type:varchar(20);comment:作者署名;NOT NULL;default:'';" json:"authorSignature"`

	CreationPurpose string `gorm:"column:creation_purpose;type:varchar(800);comment:创作目的;NOT NULL;default:'';" json:"creationPurpose"`
	CreationProcess string `gorm:"column:creation_process;type:varchar(800);comment:创作过程;NOT NULL;default:'';" json:"creationProcess"`
	Originality     string `gorm:"column:originality;type:varchar(800);comment:独创性说明;NOT NULL;default:'';" json:"originality"`
	FontCopyright   string `gorm:"column:font_copyright;type:varchar(100);comment:字体声明;NOT NULL;default:'';" json:"fontCopyright"`
	RightAuthFile   string `gorm:"column:right_auth_file;type:varchar(300);comment:肖像权授权文件;NOT NULL;default:'';" json:"rightAuthFile"`

	DciContentId      string `gorm:"column:dci_content_id;type:varchar(100);comment:dci作品ID;NOT NULL;default:'';" json:"dciContentId"`
	DigitalRegisterId string `gorm:"column:digital_register_id;type:varchar(40);comment:数登申请ID;NOT NULL;index:idx_index_digital_register_id;default:'';" json:"digitalRegisterId"`
	ApplyObtainDate   string `gorm:"column:apply_obtain_date;type:varchar(30);comment:申请发码时间;NOT NULL;default:'';" json:"ApplyObtainDate"`
	DciCodeObtainDate string `gorm:"column:dciCode_obtain_date;type:varchar(30);comment:DCI码创建时间;NOT NULL;default:'';" json:"DciCodeObtainDate"`
	//FirstCheckInfo       string                `gorm:"column:first_check_info;type:varchar(1000);comment:初审信息;NOT NULL;" json:"firstCheckInfo"`
	Remark                string `gorm:"column:remark;type:varchar(1000);comment:remark;NOT NULL;default:'';" json:"remark"`
	PreregistrationStatus string `gorm:"column:preregistration_status;type:varchar(50);comment:申领状态;NOT NULL;default:'';" json:"preregistrationStatus"`
	DigitalStatus         string `gorm:"column:digital_status;type:varchar(50);comment:数登状态;NOT NULL;default:'';" json:"digitalStatus"`
	CurrentStatusInfo     string `gorm:"column:current_status_info;type:varchar(10000);comment:当前状态说明;NOT NULL;default:'';" json:"currentStatusInfo"`
	Status                string `gorm:"column:status;type:varchar(50);comment:作品状态;NOT NULL;default:'';index:idx_status;" json:"status"`
	Source                int    `gorm:"column:source;type:tinyint(1);comment:来源1 管理后台 2 画家宝;NOT NULL;default:1;" json:"source"`
	StatusUpdateTime      int    `gorm:"column:status_update_time;bigint(20);comment:状态更新时间;NOT NULL;default:0;" json:"statusUpdateTime"`
	CreatedAt             int    `gorm:"column:created_at;type:int(11);autoCreateTime" json:"createdAt"`
	UpdatedAt             int    `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updatedAt"`
	DeletedAt             soft_delete.DeletedAt

	ArtworkInvoice   *CopyrightInvoice `gorm:"foreignKey:CopyrightId;references:Id"`
	CopyrightExtInfo *CopyrightExtInfo `gorm:"foreignKey:CopyrightId;references:Id"`
	DigitalInfo      *DigitalInfo      `gorm:"foreignKey:CopyrightId;references:Id"`
	//CopyrightLog     []*CopyrightLog   `gorm:"foreignKey:CopyrightId;references:Id"`
}
type CopyrightInvoice struct {
	CopyrightId       string                `gorm:"column:copyright_id;type:varchar(50);NOT NULL;primary_key" json:"copyrightId"`
	InvoiceType       string                `gorm:"column:invoice_type;type:varchar(20);comment:发票类型;NOT NULL;default:'';" json:"invoiceType"`
	InvoiceHeader     string                `gorm:"column:invoice_header;type:varchar(100);comment:发票抬头;NOT NULL;default:'';" json:"invoiceHeader"`
	TaxpayerNumber    string                `gorm:"column:taxpayer_number;type:varchar(100);comment:纳税人识别号;NOT NULL;default:'';" json:"taxpayerNumber"`
	RegisteredAddress string                `gorm:"column:registered_address;type:varchar(200);comment:注册地址;NOT NULL;default:'';" json:"registeredAddress"`
	RegisteredTel     string                `gorm:"column:registered_tel;type:varchar(50);comment:注册电话;NOT NULL;default:'';" json:"registeredTel"`
	OpenAccountTel    string                `gorm:"column:open_account_tel;type:varchar(50);comment:开户电话;NOT NULL;default:'';" json:"openAccountTel"`
	OpenAccountBank   string                `gorm:"column:open_account_bank;type:varchar(200);comment:开户银行;NOT NULL;default:'';" json:"openAccountBank"`
	BankAccount       string                `gorm:"column:bank_account;type:varchar(100);comment:银行账号;NOT NULL;default:'';" json:"bankAccount"`
	InvoiceCode       string                `gorm:"column:invoice_code;type:varchar(30);comment:发票代码;NOT NULL;default:'';index:inx_code;" json:"invoiceCode"`
	InvoiceNumber     string                `gorm:"column:invoice_number;type:varchar(30);comment:发票编号;NOT NULL;default:'';" json:"invoiceNumber"`
	TotalAmount       float64               `gorm:"column:total_amount;type:decimal(10,4);comment:总金额;NOT NULL;default:0.00;" json:"totalAmount"`
	OrderNo           string                `gorm:"column:order_no;type:varchar(100);comment:订单号;NOT NULL;default:'';" json:"orderNo"`
	MchNo             string                `gorm:"column:mch_no;type:varchar(100);comment:商户号;NOT NULL;default:'';" json:"mchNo"`
	CreatedAt         int                   `gorm:"column:created_at;type:int(11);autoCreateTime" json:"created_at"`
	UpdatedAt         int                   `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updated_at"`
	DeletedAt         soft_delete.DeletedAt `gorm:"column:deleted_at" json:"DeletedAt"`
}

func (CopyrightInvoice) TableName() string {
	return "copyright_invoice"
}

type CopyrightExtInfo struct {
	CopyrightId   string `gorm:"column:copyright_id;type:varchar(50);NOT NULL;primary_key" json:"id"`
	PreRegCertUrl string `gorm:"column:pre_reg_cert_url;type:varchar(300);comment:dci申领信息摘要下载地址;NOT NULL;default:'';" json:"preRegCertUrl"`

	WorkCategory string `gorm:"column:work_category;type:varchar(40);comment:作品类型;NOT NULL;default:'';" json:"workCategory"`
	FileType     string `gorm:"column:file_type;type:varchar(40);comment:文件类型;NOT NULL;default:'';" json:"fileType"`

	CreateAddress           string                `gorm:"column:create_address;type:varchar(200);comment:创作地址;NOT NULL;default:'';" json:"createAddress"`
	CreateAddressCode       string                `gorm:"column:create_address_code;type:varchar(50);comment:创作地址编码;NOT NULL;default:'';" json:"createAddressCode"`
	PublishAddress          string                `gorm:"column:publish_address;type:varchar(200);comment:发表地址;NOT NULL;default:'';" json:"publishAddress"`
	PublishAddressCode      string                `gorm:"column:publish_address_code;type:varchar(20);comment:发表地址编码;NOT NULL;default:'';" json:"publishAddressCode"`
	CreationCompletionTime  int                   `gorm:"column:creation_completion_time;type:bigint(20);comment:创作完成日期;not null;default:0;" json:"creationCompletionTime"`
	FirstPublicationTime    int                   `gorm:"column:first_publication_time;type:bigint(20);comment:首次发表日期;not null;default:0;" json:"firstPublicationTime"`
	CreatePreregisTime      int                   `gorm:"column:create_preregis_time;type:bigint(20);comment:创建申领时间;not null;default:0;" json:"createPreregisTime"`
	CreationInfo            string                `gorm:"column:creation_info;type:varchar(500);comment:创作信息;NOT NULL;default:'';" json:"creationInfo"`
	PublicationInfo         string                `gorm:"column:publication_info;type:varchar(500);comment:发表信息;NOT NULL;default:'';" json:"publicationInfo"`
	RightInfo               string                `gorm:"column:right_info;type:varchar(500);comment:权利信息;NOT NULL;default:'';" json:"rightInfo"`
	PreRegistrationTrueWill string                `gorm:"column:pre_registration_true_will;type:varchar(300);comment:真实意愿表达信息;NOT NULL;default:'';" json:"preRegistrationTrueWill"`
	CopyrightOwnerIds       string                `gorm:"column:copyright_owner_ids;type:varchar(300);comment:著作权人用户id列表;NOT NULL;default:'';" json:"copyrightOwnerIds"`
	WorkFileUrl             string                `gorm:"column:work_file_url;type:varchar(300);comment:作品文件url;NOT NULL;default:'';" json:"workFileUrl"`
	OthersWorkAuthFileId    string                `gorm:"column:others_work_auth_file_id;type:varchar(300);comment:他人作品授权文件;NOT NULL;default:'';" json:"othersWorkAuthFileId"`
	AgentName               string                `gorm:"column:agent_name;type:varchar(100);comment:代理名字;NOT NULL;default:'';" json:"agentName"`
	AgentPhone              string                `gorm:"column:agent_phone;type:varchar(100);comment:代理手机;NOT NULL;default:'';" json:"agentPhone"`
	DciJsonData             string                `gorm:"column:dci_json_data;type:text;comment:dci序列化数据;NOT NULL;" json:"dciJsonData"`
	DciCode                 string                `gorm:"column:dci_code;type:varchar(100);comment:dci编码;not null;default:'';" json:"dciCode"`
	OtherFieldList          string                `gorm:"column:other_field_list;type:varchar(4000);comment:其他文件列表;NOT NULL;default:'';" json:"otherFieldList"`
	CreatedAt               int                   `gorm:"column:created_at;type:int(11);autoCreateTime" json:"created_at"`
	UpdatedAt               int                   `gorm:"column:updated_at;type:int(11);autoCreateTime" json:"updated_at"`
	DeletedAt               soft_delete.DeletedAt `gorm:"column:deleted_at" json:"DeletedAt"`
}

func (CopyrightExtInfo) TableName() string {
	return "copyright_ext_info"
}
