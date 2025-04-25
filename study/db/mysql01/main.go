package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/exhibitionaw?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	check(err)
	//err = db.Ping()
	check(err)
	fmt.Println("Successfully connected to the database!")
	err = db.AutoMigrate(&RegisterRecord{}, &ArtworkProfile{})
	if err != nil {
		fmt.Println("Failed to migrate:", err)
	}
}

// RegisterRecord 报名记录模型
type RegisterRecord struct {
	ID                  int            `gorm:"primaryKey;autoIncrement"`                                       // 主键，自增
	CreatedAt           time.Time      `gorm:"column:created_at"`                                              // 创建时间
	UpdatedAt           time.Time      `gorm:"column:updated_at"`                                              // 更新时间
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at"`                                              // 删除时间（软删除）
	UUID                string         `gorm:"column:uuid;type:varchar(255);unique;not null;comment:报名记录Uid"`  // 唯一键
	PreliminaryRatingNo string         `gorm:"column:preliminary_rating_no;type:varchar(255);comment:初评评选号"`   // 可为空
	ReRatingNo          string         `gorm:"column:re_rating_no;type:varchar(255);comment:复评评选号"`            // 可为空
	ArtistName          string         `gorm:"column:artist_name;type:varchar(255);comment:画家姓名"`              // 可为空
	Gender              int            `gorm:"column:gender;comment:性别1男2女"`                                   // 可为空
	ArtistUUID          string         `gorm:"column:artist_uuid;type:varchar(255);comment:画家Uid"`             // 可为空
	ArtworkUUID         string         `gorm:"column:artwork_uuid;type:varchar(255);comment:作品Uid"`            // 可为空
	ArtistIDNum         string         `gorm:"column:artist_id_num;type:varchar(255);comment:画家编号"`            // 可为空
	ArtworkIDNum        string         `gorm:"column:artwork_id_num;type:varchar(255);comment:作品编号"`           // 可为空
	PhoneNum            string         `gorm:"column:phone_num;type:varchar(255);not null;comment:手机号"`        // 必填
	IDCard              string         `gorm:"column:id_card;type:varchar(255);comment:身份证号"`                  // 可为空
	Province            string         `gorm:"column:province;type:varchar(100);comment:省份"`                   // 可为空
	Address             string         `gorm:"column:address;type:varchar(3000);comment:通讯地址"`                 // 可为空
	Address1            string         `gorm:"column:address1;type:varchar(1000);comment:详细地址"`                // 可为空
	IDCardPhoto         string         `gorm:"column:id_card_photo;type:varchar(1000);comment:身份证照片"`          // 可为空
	IDCardBackPhoto     string         `gorm:"column:id_card_back_photo;type:varchar(1000);comment:身份证照片背面"`   // 可为空
	ArtistPhoto         string         `gorm:"column:artist_photo;type:varchar(1000);comment:画家本人近照"`          // 可为空
	ArtworkFile         string         `gorm:"column:artwork_file;type:varchar(1000);comment:作品文件"`            // 可为空
	ArtworkName         string         `gorm:"column:artwork_name;type:varchar(255);comment:作品名称"`             // 可为空
	ArtworkType         int            `gorm:"column:artwork_type;comment:作品类型 1 中国画"`                         // 可为空
	ArtworkSize         string         `gorm:"column:artwork_size;type:varchar(50);comment:画作尺寸"`              // 可为空
	RegisteredDate      time.Time      `gorm:"column:registered_date;comment:报名时间"`                            // 可为空
	ArtworkProfileUUID  string         `gorm:"column:artwork_profile_uuid;type:varchar(255);comment:关联作品Uuid"` // 可为空
	Artworks            ArtworkProfile `gorm:"foreignKey:ArtworkProfileUUID;references:UUID"`
}

// ArtworkProfile 画作档案模型
type ArtworkProfile struct {
	ID                         int              `gorm:"primaryKey;autoIncrement"`                                                                     // 主键，自增
	UUID                       string           `gorm:"column:uuid;type:varchar(100);unique;comment:唯一标识"`                                            // 可为空，唯一键
	Seqnum                     string           `gorm:"column:seqnum;type:varchar(256)"`                                                              // 可为空
	Tfnum                      string           `gorm:"column:tfnum;type:varchar(256)"`                                                               // 可为空
	Num                        int              `gorm:"column:num"`                                                                                   // 可为空
	ArtistName                 string           `gorm:"column:artist_name;type:varchar(256)"`                                                         // 可为空
	ArtistUUID                 string           `gorm:"column:artist_uuid;type:varchar(256)"`                                                         // 可为空
	Name                       string           `gorm:"column:name;type:varchar(256)"`                                                                // 可为空
	Belong                     int              `gorm:"column:belong"`                                                                                // 可为空
	ArtistPhoto                string           `gorm:"column:artist_photo;type:varchar(256)"`                                                        // 可为空
	PhotoPic                   string           `gorm:"column:photo_pic;type:varchar(256)"`                                                           // 可为空
	IsSign                     int              `gorm:"column:is_sign"`                                                                               // 可为空
	IsSeal                     int              `gorm:"column:is_seal"`                                                                               // 可为空
	ArtQuality                 int              `gorm:"column:art_quality"`                                                                           // 可为空
	IncompletePic              string           `gorm:"column:incomplete_pic;type:text"`                                                              // 可为空
	CopyrightPic               string           `gorm:"column:copyright_pic;type:varchar(256)"`                                                       // 可为空
	Length                     int              `gorm:"column:length"`                                                                                // 可为空
	Width                      int              `gorm:"column:width"`                                                                                 // 可为空
	Ruler                      int              `gorm:"column:ruler"`                                                                                 // 可为空
	ModelYear                  string           `gorm:"column:model_year;type:varchar(256)"`                                                          // 可为空
	ArtworkState               int              `gorm:"column:artwork_state"`                                                                         // 可为空
	ArtworkPic                 string           `gorm:"column:artwork_pic;type:varchar(256)"`                                                         // 可为空
	IsExcellent                int              `gorm:"column:is_excellent"`                                                                          // 可为空
	ScreenNum                  int              `gorm:"column:screen_num"`                                                                            // 可为空
	NetworkTrace               string           `gorm:"column:network_trace;type:varchar(256)"`                                                       // 可为空
	PhotoState                 string           `gorm:"column:photo_state;type:varchar(256)"`                                                         // 可为空
	Hash                       string           `gorm:"column:hash;type:varchar(256)"`                                                                // 可为空
	Copyright                  string           `gorm:"column:copyright;type:varchar(256)"`                                                           // 可为空
	Abstract                   string           `gorm:"column:abstract;type:varchar(2000);not null;default:'';comment:简介"`                            // 必填，默认空字符串
	Mountmode                  int              `gorm:"column:mountmode"`                                                                             // 可为空
	Material                   int              `gorm:"column:material"`                                                                              // 可为空
	Sealpic                    string           `gorm:"column:sealpic;type:varchar(256)"`                                                             // 可为空
	SealpicPhoto               string           `gorm:"column:sealpic_photo;type:varchar(600);not null;default:'';comment:手机拍摄人名章图"`                  // 必填，默认空字符串
	Signpic                    string           `gorm:"column:signpic;type:varchar(256)"`                                                             // 可为空
	InscribeDate               string           `gorm:"column:inscribe_date;type:varchar(256)"`                                                       // 可为空
	Signdate                   string           `gorm:"column:signdate;type:varchar(256)"`                                                            // 可为空
	CreatedDate                string           `gorm:"column:created_date;type:varchar(256)"`                                                        // 可为空
	CreatedAddress             string           `gorm:"column:created_address;type:varchar(256)"`                                                     // 可为空
	CreateAddressCode          string           `gorm:"column:create_address_code;type:varchar(20);not null;default:'';comment:创作地址的编码"`              // 必填，默认空字符串
	CopyrightCreateAddress     string           `gorm:"column:copyright_create_address;type:varchar(100);not null;default:'';comment:'版权创建地址'"`       // 必填，默认空字符串
	CopyrightCreateAddressCode string           `gorm:"column:copyright_create_address_code;type:varchar(20);not null;default:'';comment:'版权创建地址编码'"` // 必填，默认空字符串
	ArriveTime                 string           `gorm:"column:arrive_time;type:varchar(256)"`                                                         // 可为空
	ArtworkType                int              `gorm:"column:artwork_type"`                                                                          // 可为空
	GiftInfo                   string           `gorm:"column:gift_info;type:varchar(256)"`                                                           // 可为空
	Scroll                     string           `gorm:"column:scroll;type:varchar(256)"`                                                              // 可为空
	Comment                    string           `gorm:"column:comment;type:varchar(2000);not null;default:'';comment:画作备注"`                           // 必填，默认空字符串
	ArtMeansOfExpression       string           `gorm:"column:art_means_of_expression;type:varchar(256)"`                                             // 可为空
	Size                       int              `gorm:"column:size"`                                                                                  // 可为空
	ArtHorizontal              int              `gorm:"column:art_horizontal"`                                                                        // 可为空
	FlowState                  int              `gorm:"column:flow_state"`                                                                            // 可为空
	HDPic                      string           `gorm:"column:hd_pic;type:varchar(256)"`                                                              // 可为空
	ArtCondition               int              `gorm:"column:art_condition"`                                                                         // 可为空
	Status                     int              `gorm:"column:status"`                                                                                // 可为空
	PriceRuler                 float32          `gorm:"column:price_ruler"`                                                                           // 可为空
	PriceCopyright             float32          `gorm:"column:price_copyright"`                                                                       // 可为空
	PriceArtwork               float32          `gorm:"column:price_artwork"`                                                                         // 可为空
	PriceMarket                float32          `gorm:"column:price_market"`                                                                          // 可为空
	Rate                       string           `gorm:"column:rate;type:varchar(100);not null;default:'';comment:比率"`                                 // 必填，默认空字符串
	PriceRun                   float64          `gorm:"column:price_run;type:decimal(10,2);not null;default:0.00;comment:润格"`                         // 必填，默认 0.00
	FilterState                int              `gorm:"column:filter_state;not null;default:1;comment:筛选状态 1 通过 2 不通过"`                               // 必填，默认 1
	CreateSource               int              `gorm:"column:create_source;not null;default:1;comment:来源 1 后台 2 画家宝"`                                // 必填，默认 1
	CreateDoneDate             string           `gorm:"column:create_done_date;type:varchar(100);not null;default:'';comment:创作完成时间"`                 // 必填，默认空字符串
	Mask                       int8             `gorm:"column:mask;type:tinyint;not null;default:1;comment:画作标记 1 一手画 2 二手画"`                         // 必填，默认 1
	InSource                   int8             `gorm:"column:in_source;type:tinyint;not null;default:1;comment:哪个系统 1 管理 2 画家宝 3 管理和画家宝"`            // 必填，默认 1
	CreatedAt                  time.Time        `gorm:"column:created_at"`                                                                            // 可为空
	UpdatedAt                  time.Time        `gorm:"column:updated_at"`                                                                            // 可为空
	DeletedAt                  gorm.DeletedAt   `gorm:"column:deleted_at;type:bigint unsigned"`
	Records                    []RegisterRecord `gorm:"foreignKey:ArtworkProfileUUID;references:UUID"` // 可为空
}
