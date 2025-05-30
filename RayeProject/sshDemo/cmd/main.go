package main

import (
	"context"
	"errors"
	"fmt"
	"ssh/demo/api/dci"
	"ssh/demo/config"
	"ssh/demo/model"
	"ssh/demo/model/digi"
	"ssh/demo/pkg/db"
	"ssh/demo/pkg/service"
	sshconn "ssh/demo/pkg/sshConn"
	"ssh/demo/pkg/utils"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	failedRecoed  = "../data/failed_uuids-5-23.txt"
	successRecoed = "../data/success_uuids-5-23.txt"
)

func init() {
	err := config.InitConfig()
	if err != nil {
		zap.L().Error("config.InitConfig() failed", zap.Error(err))
		return
	}
}
func main() {
	service.Ser()
	sshClient, err := sshconn.SshConnect()
	if err != nil {
		return
	}
	defer sshClient.Close()

	// 使用封装后的函数
	localPort, stopChan, err := sshconn.StartTunnelForwarding(sshClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("本地隧道端口: %d\n", localPort)

	// 等待隧道准备就绪
	time.Sleep(1 * time.Second)
	Db, err := db.Init(localPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	DbArtwork := Db[db.Artwork]
	DbArtist := Db[db.Artist]
	DbDigiCopy := Db[db.DigiCopy]
	//fmt.Println("Successfully connected to MySQL via SSH!")
	node := utils.NewSf()
	var uuids []string
	randId := node.Generate().String()
	var total int64
	data := []string{
		"T10821001-3",
	}
	var successRecords []model.ResultRecord
	var failedRecords []model.ResultRecord
	//1为一手画
	//找到版权登记的一手画所有uid
	fmt.Println(DbArtwork == nil)
	DbArtwork.Table("artwork_profile").
		//Select("artwork_copyright.artwork_uuid").
		Select("uuid").
		//Joins("JOIN artwork_copyright ON artwork_profile.uuid = artwork_copyright.artwork_uuid").
		//	Where("artwork_copyright.serial_number != ? and artwork_copyright.cert_digi !=?", "", "").
		Where("tfnum in (?)", data).
		Find(&uuids)
	fmt.Println("水水水水")
	fmt.Println("uuids:", len(uuids))
	var uuidss []string
	for i := 0; i < len(uuids); i++ {
		// if total == 500 {
		// 	break
		// }
		count := int64(0)
		res := make(map[string]interface{})
		DbDigiCopy.Table("artwork_copyright").Where("artwork_uuid  =?", uuids[i]).Count(&count).Take(&res)
		//fmt.Println("res: ", res["artist_uuid"])
		if count >= 1 {
			fmt.Println(count)
			fmt.Println("查询已存在的作品: ", uuids[i])
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:        uuids[i],
				Message:     "作品已存在",
				OperateDate: time.Now().Format("2006-01-02 15:04:05"),
			})

			// var dciUserExists bool
			// if err := DbDigiCopy.Table("dci_user").Select("count(*) > 0").
			// 	Where("artist_uuid = ?", res["artist_uuid"]).
			// 	Find(&dciUserExists).Error; err != nil {
			// 	fmt.Printf("查询dci_user表失败: %v\n", err)
			// 	continue
			// }
			// fmt.Println("dhsghv:", dciUserExists)
			// if dciUserExists {
			// 	continue
			// }
			//fmt.Println("count=1: ", uuids[i])
			continue
		}
		uuidss = append(uuidss, uuids[i])
		fmt.Println("*********************start: *********************", i)
		fmt.Println("uuid: ", uuids[i])
		total++
		artwork_profile := make(map[string]interface{})
		artwork_ext_data := make(map[string]interface{})
		artist := make(map[string]interface{})
		artwork_copyright := make(map[string]interface{})
		artist_ext_data := make(map[string]interface{})
		DbArtwork.Table("artwork_profile").Where("uuid=?", uuids[i]).Take(artwork_profile)
		DbArtwork.Table("artwork_ext_data").Where("artwork_uuid=?", uuids[i]).Take(artwork_ext_data)
		DbArtwork.Table("artwork_copyright").Where("artwork_uuid=?", uuids[i]).Take(artwork_copyright)
		DbArtist.Table("artist_profile").Where("uid=?", artwork_profile["artist_uuid"]).Take(artist)
		DbArtist.Table("artist_ext_data").Where("uid=?", artwork_profile["artist_uuid"]).Take(artist_ext_data)
		// fmt.Println(artwork_profile["artist_name"], artist["tnum"], artwork_profile["name"], artwork_profile["tfnum"], "作品类型", "原创字体", artwork_profile["copyright_create_address"],
		// 	artwork_ext_data["creation_purpose"], artwork_ext_data["creation_process"], artwork_ext_data["originality"], "委托书：", artwork_copyright["promise_letter_url"], "承诺书：",
		// 	artwork_copyright["entrust_letter_url"], artwork_ext_data["publish_status"], artwork_ext_data["publish_address_name"], artwork_ext_data["publish_date"], "作者署名", "作品性质", artwork_profile["create_done_date"],
		// 	"全部权利", artwork_copyright["apply_time"], artwork_copyright["serial_number"], artwork_copyright["cert-register_time"], artwork_copyright["cert_digi"])
		// fmt.Println()
		//fmt.Println(randId)

		digi_copyright := digi.ArtworkCopyright{
			Id:              randId, //artwork_copyright["serial_number"].(string),
			ArtistUuid:      artist["uid"].(string),
			ArtworkUuid:     artwork_profile["uuid"].(string),
			AuthorName:      artwork_profile["artist_name"].(string),
			Tfnum:           artwork_profile["tfnum"].(string),
			WorkName:        artwork_profile["name"].(string),
			FontCopyright:   "ORIGINAL_FONT",
			CreationPurpose: artwork_ext_data["creation_purpose"].(string),
			CreationProcess: artwork_ext_data["creation_process"].(string),
			Originality:     artwork_ext_data["originality"].(string),
			AuthorSignature: artwork_profile["artist_name"].(string),
			Status:          "FINISH",
			CreatedAt:       int(artwork_profile["created_at"].(int32)),
			UpdatedAt:       int(time.Now().Unix()),
			Source:          3,
		}
		var certRegisterTime int
		if datestr, ok := artwork_copyright["cert_register_time"].(string); ok {
			if datestr == "" {
				certRegisterTime = int(time.Now().Unix())
			} else {
				// 尝试解析为日期字符串
				t, err := time.Parse("2006-01-02", datestr)
				if err != nil {
					// 如果不是日期格式，尝试解析为Unix时间戳字符串
					if ts, err := strconv.ParseInt(datestr, 10, 64); err == nil {
						certRegisterTime = int(ts)
					} else {
						fmt.Println("解析cert_register_time失败:", err)
						failedRecords = append(failedRecords, model.ResultRecord{
							UUID:    uuids[i],
							Message: fmt.Sprintf("转换cert_register_time失败: 既不是日期格式也不是时间戳格式"),
						})
						continue
					}
				} else {
					certRegisterTime = int(t.Unix())
				}
			}
		}
		if certRegisterTime != 0 {
			digi_copyright.StatusUpdateTime = certRegisterTime
		} else {
			digi_copyright.StatusUpdateTime = int(time.Now().Unix())
		}
		var creationCompletionTime int
		if dateStr, ok := artwork_profile["create_done_date"].(string); ok {
			if dateStr == "" {
				creationCompletionTime = int(time.Now().Unix())
			} else {
				// 解析日期字符串
				t, err := time.Parse("2006-01-02", dateStr)
				if err != nil {
					fmt.Println("解析create_done_date失败:", err)
					failedRecords = append(failedRecords, model.ResultRecord{
						UUID:    uuids[i],
						Message: fmt.Sprintf("解析create_done_date失败: %v", err),
					})
					continue
				}
				creationCompletionTime = int(t.Unix())
			}
		}
		digi_copy_ext_info := digi.CopyrightExtInfo{
			CopyrightId:          randId, //artwork_copyright["serial_number"].(string),
			WorkCategory:         "ART",
			CreateAddress:        artwork_profile["copyright_create_address"].(string),
			RightInfo:            `{"rightscope":"ALL","rightObtainWay":"ORIGINAL","ownershipWay":"PERSON"}`,
			WorkFileUrl:          artwork_ext_data["digi_art_img"].(string),
			OthersWorkAuthFileId: artwork_copyright["entrust_letter_url"].(string),
			PublishAddress:       artwork_ext_data["publish_address_name"].(string),
			PublishAddressCode:   artwork_ext_data["publish_address_code"].(string),
		}
		publishDate, ok, err := utils.ConvertFieldToInt(artist_ext_data["publish_date"], "publish_date")
		if err != nil {
			fmt.Println("转换publish_date失败:", err)
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("转换publish_date失败: %v", err),
			})
			continue
		}
		if ok {
			digi_copy_ext_info.FirstPublicationTime = publishDate
		} else {
			digi_copy_ext_info.FirstPublicationTime = int(time.Now().Unix())
		}
		if digi_copy_ext_info.CreateAddress == "" {
			digi_copy_ext_info.CreateAddress = "江苏省,苏州市"
			digi_copy_ext_info.CreateAddressCode = "320500"
		}
		if digi_copy_ext_info.CreationCompletionTime == 0 {
			digi_copy_ext_info.CreationCompletionTime = int(time.Now().Unix())
		}
		if digi_copy_ext_info.PublishAddress == "" {
			digi_copy_ext_info.PublishAddress = "江苏省,苏州市"
			digi_copy_ext_info.PublishAddressCode = "320500"
			digi_copy_ext_info.FirstPublicationTime = int(time.Now().Unix())
		}
		fmt.Println("digi_art: ", digi_copy_ext_info.WorkFileUrl)
		if ok {
			digi_copy_ext_info.CreationCompletionTime = creationCompletionTime
		} else {
			digi_copy_ext_info.CreationCompletionTime = int(time.Now().Unix())
		}
		digi_copy_invoice := digi.CopyrightInvoice{
			CopyrightId: randId, //artwork_copyright["serial_number"].(string),
		}
		digi_digi_info := digi.DigitalInfo{
			Id:                     randId,
			CopyrightId:            artwork_copyright["serial_number"].(string),
			DigitalRegisterCertUrl: artwork_copyright["cert_digi"].(string),
			RegNumber:              artwork_copyright["register_number"].(string),
		}
		var applyTime int
		if date, ok := artwork_copyright["apply_time"].(string); ok {
			if date == "" {
				applyTime = int(time.Now().Unix())
			} else {
				// 尝试解析为日期字符串
				t, err := time.Parse("2006-01-02", date)
				if err != nil {
					// 如果不是日期格式，尝试解析为Unix时间戳字符串
					if ts, err := strconv.ParseInt(date, 10, 64); err == nil {
						certRegisterTime = int(ts)
					} else {
						fmt.Println("解析cert_register_time失败:", err)
						failedRecords = append(failedRecords, model.ResultRecord{
							UUID:    uuids[i],
							Message: fmt.Sprintf("转换cert_register_time失败: 既不是日期格式也不是时间戳格式"),
						})
						continue
					}
				} else {
					applyTime = int(t.Unix())
				}
			}
		}
		if err != nil {
			fmt.Println("转换apply_time失败:", err)
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("转换apply_time失败: %v", err),
			})
			continue
		}
		if ok {
			digi_digi_info.DigitalRegisterApplyTime = applyTime
		}

		queryUserReq := dci.QueryDciUserRequest{
			CertificateType:   "IDENTITY_CARD",
			CertificateNumber: artist["card_id"].(string),
			Phone:             strings.TrimSpace(artist["phone"].(string)),
		}
		fmt.Println("phone: ", queryUserReq.Phone)
		var queryDciUserResp *dci.QueryDciUserResponse
		if queryDciUserResp, err = service.GrpcDciImpl.QueryDciUser(context.Background(), &queryUserReq); err != nil {
			fmt.Println("CheckDciReg err queryUserReq", zap.Error(err))
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("CheckDciReg err queryUserReq: %v", err),
			})
			continue
		}
		dciUser := digi.DciUser{}
		fmt.Println("开始查询dciUser:", i)
		if queryDciUserResp.ResultCode == "OK" {
		} else {
			if queryDciUserResp.ResultCode == "BAD_REQUEST" && queryDciUserResp.ResultMsg != "DCI_USER_NOT_EXIST" {
				failedRecords = append(failedRecords, model.ResultRecord{
					UUID:    uuids[i],
					Message: fmt.Sprintf("BAD_REQUEST: %v", err),
				})
				continue
			}
			switch queryDciUserResp.ResultMsg {
			//case "DCI_USER_ALREADY_EXIST":
			case "DCI_USER_NOT_EXIST":
				if artist_ext_data["card_face"].(string) == "" || artist_ext_data["card_national"].(string) == "" {
					failedRecords = append(failedRecords, model.ResultRecord{
						UUID:    uuids[i],
						Message: fmt.Sprintf("身份证信息不全: %v", err),
					})
					continue
				}
				dciCardFaceFileId, errF := service.GrpcDciImpl.GetUploadUrl(context.Background(), &dci.GetUploadUrlRequest{FileName: artist_ext_data["card_face"].(string)})
				dciCardNationalFileId, errB := service.GrpcDciImpl.GetUploadUrl(context.Background(), &dci.GetUploadUrlRequest{FileName: artist_ext_data["card_national"].(string)})
				if errF != nil || errB != nil {
					err = errors.New("上传错误")
					failedRecords = append(failedRecords, model.ResultRecord{
						UUID:    uuids[i],
						Message: fmt.Sprintf("上传错误: %v", err),
					})

					continue
				}
				dciUserAddReq := &dci.AddDciUserRequest{
					CertName:               artist["name"].(string),
					CertificateNumber:      artist["card_id"].(string),
					CertificateType:        "IDENTITY_CARD",
					CertificateFrontFileId: dciCardFaceFileId.FileId,
					CertificateBackFileId:  dciCardNationalFileId.FileId,
					Phone:                  strings.TrimSpace(artist["phone"].(string)),
					AreaType:               "CHINA_MAINLAND",
				}
				dciAddUserResp, _ := service.GrpcDciImpl.AddDciUser(context.Background(), dciUserAddReq)
				if dciAddUserResp == nil {
					err = errors.New("数据为空")
					failedRecords = append(failedRecords, model.ResultRecord{
						UUID:    uuids[i],
						Message: fmt.Sprintf("数据为空: %v", err),
					})

					continue
				}
				if dciAddUserResp.ResultCode != "OK" {
					err = errors.New(dciAddUserResp.ResultMsg)
					failedRecords = append(failedRecords, model.ResultRecord{
						UUID:    uuids[i],
						Message: fmt.Sprintf("`dciAddUserResp.ResultCode != OK`: %v", err),
					})
					continue
				}
				if dciAddUserResp.DciUserStatus != "NORMAL" {
					err = errors.New("用户已停用")
					failedRecords = append(failedRecords, model.ResultRecord{
						UUID:    uuids[i],
						Message: fmt.Sprintf("用户已停用: %v", err),
					})
					continue
				}
				DciUserId := dciAddUserResp.DciUserId
				dciUserStatus := dciAddUserResp.DciUserStatus
				//queryDciUserResp.DciUserId = dciUserId
				if err := DbDigiCopy.Model(&digi.DciUser{}).Where("artist_uuid=?", artist["uid"]).First(&dciUser).Error; err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						fmt.Println("未注册") // 没有注册
						dciUserId := node.Generate().Int64()
						dciUser = digi.DciUser{
							Id:                dciUserId,
							ArtistUuid:        artwork_profile["artist_uuid"].(string),
							CertName:          artist["name"].(string),
							CertificateNumber: artist["card_id"].(string),
							CertificateType:   "IDENTITY_CARD",
							CardFaceUrl:       artist_ext_data["card_face"].(string),
							CardNationUrl:     artist_ext_data["card_national"].(string),
							Phone:             artist["phone"].(string),
							AreaType:          "CHINA_MAINLAND",
							DciUserStatus:     dciUserStatus,
							DciUserId:         DciUserId,
						}
						if dciUser.CertificateNumber == "" {
							failedRecords = append(failedRecords, model.ResultRecord{
								UUID:    uuids[i],
								Message: fmt.Sprintf("身份证号码为空: %v", err),
							})
							continue
						}
						if dciUser.CertName == "" {
							failedRecords = append(failedRecords, model.ResultRecord{
								UUID:    uuids[i],
								Message: fmt.Sprintf("姓名为空: %v", err),
							})
						}
						if dciUser.CardFaceUrl == "" {
							failedRecords = append(failedRecords, model.ResultRecord{
								UUID:    uuids[i],
								Message: fmt.Sprintf("身份证正面为空: %v", err),
							})
							continue
						}
						if dciUser.CardNationUrl == "" {
							failedRecords = append(failedRecords, model.ResultRecord{
								UUID:    uuids[i],
								Message: fmt.Sprintf("身份证反面为空: %v", err),
							})
							continue
						}

						if err = DbDigiCopy.Model(&digi.DciUser{}).Create(&dciUser).Error; err != nil {
							fmt.Println("dciuser create err:", err)
							failedRecords = append(failedRecords, model.ResultRecord{
								UUID:    uuids[i],
								Message: fmt.Sprintf("dciuser create err: %v", err),
							})
							continue
						}
					} else {
						fmt.Println("dciuser query err:", err)
						failedRecords = append(failedRecords, model.ResultRecord{
							UUID:    uuids[i],
							Message: fmt.Sprintf("dciuser query err: %v", err),
						})
						continue
					}
				} else {
					dciUser.DciUserStatus = dciUserStatus
					dciUser.DciUserId = DciUserId
					if err = DbDigiCopy.Model(&digi.DciUser{}).Where("artist_uuid=?", artist["uid"]).Updates(&dciUser).Error; err != nil {
						fmt.Println("dciuser update err:", err)
						failedRecords = append(failedRecords, model.ResultRecord{
							UUID:    uuids[i],
							Message: fmt.Sprintf("dciuser update err: %v", err),
						})
						continue
					}
				}
			default:
				err = errors.New(queryDciUserResp.ResultMsg)
				failedRecords = append(failedRecords, model.ResultRecord{
					UUID:    uuids[i],
					Message: fmt.Sprintf("default: %v", err),
				})
				return
			}
		}
		digi_copyright.DciUserId = dciUser.DciUserId
		uploadReq := &dci.GetUploadUrlRequest{
			FileName: artwork_ext_data["digi_art_img"].(string),
		}
		dciUploadResp, err := service.GrpcDciImpl.GetUploadUrl(context.Background(), uploadReq)
		if err != nil {
			zap.L().Info("DciUploadUrl", zap.Any("dciUploadResp err", err.Error()))
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("dciUploadResp err: %v", err),
			})
			fmt.Println("错误：", err)
			//continue
		}
		if queryDciUserResp != nil && dciUploadResp.ResultCode != "OK" {
			err = errors.New(dciUploadResp.ResultMsg)
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("1111: %v", err),
			})

		}
		//digi_copy_ext_info.WorkFileUrl = dciUploadResp.FileId
		// fmt.Println("dci: ", queryDciUserResp)
		fmt.Println("********")
		//fmt.Println("***************************************************************")
		tx := DbDigiCopy.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		// 创建主表记录
		if err := tx.Model(&digi.ArtworkCopyright{}).Create(&digi_copyright).Error; err != nil {
			fmt.Println("创建ArtworkCopyright失败:", err)
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("创建ArtworkCopyright失败: %v", err),
			})
			continue
		}
		fmt.Println("over*")
		// 创建关联表记录

		if err := tx.Model(digi.CopyrightExtInfo{}).Create(&digi_copy_ext_info).Error; err != nil {
			fmt.Println("创建CopyrightExtInfo失败:", err)
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("创建CopyrightExtInfo失败: %v", err),
			})
			continue
		}
		if err := tx.Model(&digi.CopyrightInvoice{}).Create(&digi_copy_invoice).Error; err != nil {
			fmt.Println("创建CopyrightInvoice失败:", err)
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("创建CopyrightInvoice失败: %v", err),
			})
			continue
		}

		if err := tx.Create(&digi_digi_info).Error; err != nil {
			fmt.Println("创建DigitalInfo失败:", err)
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:    uuids[i],
				Message: fmt.Sprintf("创建DigitalInfo失败: %v", err),
			})
			continue
		}
		// 提交事务
		if err := tx.Commit().Error; err != nil {
			fmt.Println("提交事务失败:", err)
			failedRecords = append(failedRecords, model.ResultRecord{
				UUID:        uuids[i],
				Message:     fmt.Sprintf("提交事务失败: %v", err),
				OperateDate: time.Now().Format("2006-01-02 15:04:05"),
			})
			continue
		}
		successRecords = append(successRecords, model.ResultRecord{
			UUID:        uuids[i],
			Message:     "ok",
			OperateDate: time.Now().Format("2006-01-02 15:04:05"),
		})
		fmt.Println(uuids[i])
	}

	// 使用新的写入方法
	if err := utils.WriteRecordsToFile(failedRecords, failedRecoed); err != nil {
		failedRecords = append(failedRecords, model.ResultRecord{
			UUID:        "",
			Message:     fmt.Sprintf("写入失败记录文件错误: %v", err),
			OperateDate: time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	if err := utils.WriteRecordsToFile(successRecords, successRecoed); err != nil {
		failedRecords = append(failedRecords, model.ResultRecord{
			UUID:    "",
			Message: fmt.Sprintf("写入成功记录文件错误: %v", err),
		})
	}

	fmt.Println()
	fmt.Println("total", total)

	//DbDigiCopy.Table("artwork_copyright").Select("artwork_uuid").Where("artwork_uuid  IN (?)", uuids).Find(&res)
	fmt.Println("uuids", len(uuids))

	//fmt.Println(uuidss[1:4])
	// 停止隧道
	close(stopChan)
}
