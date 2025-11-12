package fiee

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"test001/model"
	"time"

	"golang.org/x/crypto/ssh"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func UpdateVideoTime() {
	// SSH 配置
	sshConfig := &ssh.ClientConfig{
		User: "",
		Auth: []ssh.AuthMethod{
			ssh.Password(""), // 也可以用 ssh.PublicKeys(...)
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 测试用，生产要换安全方式
		Timeout:         time.Second * 5,
	}

	// 建立 SSH 连接
	sshClient, err := ssh.Dial("tcp", ":22", sshConfig)
	if err != nil {
		panic("SSH连接失败: " + err.Error())
	}

	// 创建 MySQL 的远程连接转发
	mysqlHost := ":3306" // 远程数据库地址
	localPort := "3307"  // 本地监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:"+localPort)
	if err != nil {
		panic("本地监听端口失败: " + err.Error())
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go func() {
				remote, err := sshClient.Dial("tcp", mysqlHost)
				if err != nil {
					fmt.Println("转发失败:", err)
					return
				}
				go copyConn(conn, remote)
				go copyConn(remote, conn)
			}()
		}
	}()

	// 等待转发生效
	time.Sleep(time.Second * 1)

	// 使用 GORM 连接本地转发端口
	user := ""
	pwd := ""
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/micro_bundle?charset=utf8mb4&parseTime=True&loc=Local", user, pwd)
	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	sqlDB, _ := db1.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	db1 = db1.Debug()
	// dsn2 := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/micro-account?charset=utf8mb4&parseTime=True&loc=Local", user, pwd)
	// db2, err := gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	// if err != nil {
	// 	panic("数据库连接失败: " + err.Error())
	// }

	adjustTimes(db1)

}

type ArtistPay struct {
	CustomerId string
	PayTime    string
}

func adjustTimes(db *gorm.DB) {
	var uuids []string
	var errors []string

	var artistPayTime = make(map[string]string)
	var artistPay []ArtistPay
	db.Model(&model.BundleOrderRecords{}).Select("customer_id", " pay_time").Where("deleted_at is null").Scan(&artistPay)
	for _, v := range artistPay {
		artistPayTime[v.CustomerId] = v.PayTime
	}
	//artistUid := []string{"428", "432", "433", "425"}
	artistUid := []string{}
	// 先取出所有符合条件的 uuid
	if err := db.Model(&model.Work{}).
		Where("status = ? AND source = ? AND deleted_at = 0   and artist_uuid not in (?)", 6, 2, artistUid).
		//Where("artist_uuid = ? AND deleted_at = 0", 165).
		Pluck("uuid", &uuids).Error; err != nil {
		errors = append(errors, fmt.Sprintf("获取UUID失败: %v", err))
		writeErrors(errors)
		return
	}
	fmt.Println(len(uuids))

	// uuids = []string{"5cef7b21-c885-4674-a4dd-b23ba51fc2d8"}
	i := 0
	//checked := make(map[string]bool)

	for _, v := range uuids {
		i++

		// tx := db.Begin()
		// if tx.Error != nil {
		// 	errors = append(errors, fmt.Sprintf("开启事务失败 uuid=%s: %v", v, tx.Error))
		// 	continue
		// }
		tx := db
		type result struct {
			WorkUuid   string
			ArtistUuid string
			CaoGaoTime string
		}
		var res result
		if err := tx.Model(&model.WorkLog{}).
			Select("work_uuid,artist_uuid,update_time").
			Where("work_uuid = ? AND work_status = ? AND deleted_at = 0", v, 1).
			Order("update_time DESC").
			Limit(1).
			Scan(&res).Error; err != nil {
			fmt.Println(err)
			continue
		}
		// var publishTimeStr string
		// // 查出对应的 status=6 的 update_time
		// if err := tx.Model(&model.WorkLog{}).
		// 	Select("update_time").
		// 	Where("work_uuid = ? AND work_status = ? AND deleted_at = 0", v, 6).
		// 	Order("update_time DESC").
		// 	Limit(1).
		// 	Scan(&publishTimeStr).Error; err != nil {
		// 	errors = append(errors, fmt.Sprintf("获取 publishTime 失败 work_uuid=%s: %v", v, err))
		// 	fmt.Println(err)
		// 	continue
		// }
		payTimeStr := ""
		if t, ok := artistPayTime[res.ArtistUuid]; ok {
			payTimeStr = t
		} else {
			errors = append(errors, fmt.Sprintf("获取 payTime 失败 work_uuid=%s: %v", v, "没有找到"))
			continue
		}
		// 转换时间
		payTime, err := time.Parse("2006-01-02 15:04:05", payTimeStr)
		if err != nil {
			errors = append(errors, fmt.Sprintf("时间解析失败 work_uuid=%s: %s (%v)", v, payTimeStr, err))
			fmt.Println(err)
			continue
		}
		caoGaoTime, err := time.Parse("2006-01-02 15:04:05", res.CaoGaoTime)
		if err != nil {
			errors = append(errors, fmt.Sprintf("时间解析失败 work_uuid=%s: %s (%v)", v, res.CaoGaoTime, err))
			fmt.Println(err)
			continue
		}
		// p := publishTime.Add(-120 * time.Minute)
		// if payTime.Before(p) {
		// 	continue
		// }
		//if payTime.Month() == 7 || payTime.Month() == 6 || payTime.Month() <= 5 {
		//publishTime = adjustPublishTime(payTime, publishTime)
		//}

		submitTime, _, _, err := generateRandomTimes(payTime, caoGaoTime)
		if err != nil {
			errors = append(errors, fmt.Sprintf("时间生成失败 work_uuid=%s: %s (%v)", v, res.CaoGaoTime, err))
			fmt.Println(err)
			continue
		}
		// fmt.Println(caoGaoTime, submitTime)
		// if err := tx.Model(&model.WorkLog{}).
		// 	Where("work_uuid = ? AND work_status = ?", v, 6).
		// 	UpdateColumn("update_time", publishTime.Format("2006-01-02 15:04:05")).Error; err != nil {
		// 	errors = append(errors, fmt.Sprintf("更新 WorkLog 发布成功时间失败 work_uuid=%s: %v", v, err))
		// 	fmt.Println(err)
		// 	continue
		// }
		// fmt.Println("publishTime 更新成功")

		// 更新 WorkLog status=1 的时间
		// if err := tx.Model(&model.WorkLog{}).
		// 	Where("work_uuid = ? AND work_status = ?", v, 1).
		// 	UpdateColumn("update_time", caogaoTime.Format("2006-01-02 15:04:05")).Error; err != nil {
		// 	errors = append(errors, fmt.Sprintf("更新 WorkLog 草稿时间失败 work_uuid=%s: %v", v, err))
		// 	fmt.Println(err)
		// 	continue
		// }
		//fmt.Println("更新 WorkLog 草稿时间成功")

		// 更新 Work 的 submit_time
		if err := tx.Model(&model.Work{}).
			Where("uuid = ?", v).
			UpdateColumn("submit_time", submitTime.Format("2006-01-02 15:04:05")).Error; err != nil {
			errors = append(errors, fmt.Sprintf("更新 Work 提交时间失败 uuid=%s: %v", v, err))
			fmt.Println(err)
			continue
		}
		fmt.Println("更新 Work 提交时间成功")
		// // 更新 CostLog 的 submit_time 为 publishTime
		// if err := tx.Model(&model.CostLog{}).
		// 	Where("work_uuid = ?", v).
		// 	UpdateColumn("submit_time", publishTime.Format("2006-01-02 15:04:05")).Error; err != nil {
		// 	errors = append(errors, fmt.Sprintf("更新 CostLog 时间失败 work_uuid=%s: %v", v, err))
		// 	fmt.Println(err)
		// 	continue
		// }
		// fmt.Println("extra 更新")

		// 需要更新的字段
		// updates := model.WorkExtra{
		// 	WorkUuid:            v,
		// 	ArtistConfirmedTime: publishTime.In(time.FixedZone("CST", 8*3600)).Unix(),
		// 	CreatedAt:           int(publishTime.In(time.FixedZone("CST", 8*3600)).Unix()),
		// 	UpdatedAt:           int(publishTime.In(time.FixedZone("CST", 8*3600)).Unix()),
		// }

		// if err := tx.Clauses(clause.OnConflict{
		// 	Columns:   []clause.Column{{Name: "work_uuid"}}, // 唯一键
		// 	DoUpdates: clause.AssignmentColumns([]string{"artist_confirmed_time", "created_at", "updated_at"}),
		// }).Create(&updates).Error; err != nil {
		// 	errors = append(errors, fmt.Sprintf("Upsert 失败 work_uuid=%s: %v", v, err))
		// 	fmt.Println(err)
		// 	continue
		// }

		// fmt.Println("更新 workExta 时间成功")

		// 提交事务
		// if err := tx.Commit().Error; err != nil {
		// 	errors = append(errors, fmt.Sprintf("提交事务失败 uuid=%s: %v", v, err))
		// 	tx.Rollback()
		// 	continue
		// }
		//checked[v] = true

		// if err := encoder.Encode(record); err != nil {
		// 	log.Printf("写入 JSONL 失败: %v", err)
		// }
		fmt.Println("***********************", i, "/", len(uuids))
	}

	// 写错误日志
	writeErrors(errors)
}
func writeErrors(errors []string) {
	if len(errors) == 0 {
		log.Println("✅ 所有数据处理完成，没有错误。")
		return
	}

	file, err := os.Create("errors.txt")
	if err != nil {
		log.Fatalf("❌ 创建错误日志文件失败: %v", err)
	}
	defer file.Close()

	for _, e := range errors {
		file.WriteString(e + "\n")
	}

	log.Printf("⚠️ 处理完成，有 %d 条错误，已写入 errors.txt\n", len(errors))
}
func adjustPublishTime(payTime, publishTime time.Time) time.Time {
	loc := payTime.Location()
	year, month, _ := payTime.Date()
	month = publishTime.Month()
	// 计算当月有多少天
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	nextMonth := firstOfMonth.AddDate(0, 1, 0)
	daysInMonth := nextMonth.Add(-time.Hour).Day()

	for {
		// 随机日、时、分、秒
		day := rand.Intn(daysInMonth) + 1
		hour := rand.Intn(24)
		min := rand.Intn(60)
		sec := rand.Intn(60)

		candidate := time.Date(year, month, day, hour, min, sec, 0, loc)

		// 条件：
		// 1. publishTime 必须在 payTime 之后至少 1h5min
		// 2. 同一个月
		if candidate.After(payTime.Add(120*time.Minute)) && candidate.Month() == month {
			return candidate
		}
	}
}
func generateRandomTimes(payTime, caoGaoTime time.Time) (time.Time, time.Time, time.Time, error) {
	// 验证时间顺序
	rand.Seed(time.Now().UnixNano())

	adjustedCaoGaoTime := adjustCaoGaoTime(payTime, caoGaoTime)
	minInterval := 5 * time.Minute
	submitTime := generateSubmitTime(payTime, adjustedCaoGaoTime, minInterval)
	publishTime := generatePublishTime(adjustedCaoGaoTime)

	return submitTime, adjustedCaoGaoTime, publishTime, nil
}

// 调整 caoGaoTime，避免在跨月边界（距离月末不超过2天就往前提）
func adjustCaoGaoTime(payTime, caoGaoTime time.Time) time.Time {
	// 获取 caoGaoTime 所在月份的最后一天
	lastDayOfMonth := time.Date(caoGaoTime.Year(), caoGaoTime.Month()+1, 0, 23, 59, 59, 0, caoGaoTime.Location())
	timeToMonthEnd := lastDayOfMonth.Sub(caoGaoTime)

	if timeToMonthEnd <= 48*time.Hour {
		advanceTime := 48*time.Hour - timeToMonthEnd + time.Hour

		newCaoGaoTime := caoGaoTime.Add(-advanceTime)

		if !newCaoGaoTime.After(payTime) {
			newCaoGaoTime = generateRandomTimeAfter(payTime)
		}

		return newCaoGaoTime
	}

	return caoGaoTime
}

// 生成 submitTime
func generateSubmitTime(payTime, caoGaoTime time.Time, minInterval time.Duration) time.Time {
	totalDuration := caoGaoTime.Sub(payTime)

	if totalDuration < 2*time.Hour {
		submitOffset := time.Duration(rand.Int63n(int64(totalDuration-minInterval))) + minInterval
		return payTime.Add(submitOffset)
	}

	// 计算可用的提交时间范围（在 payTime 之后，caoGaoTime 之前）
	minSubmitTime := payTime.Add(minInterval)
	maxSubmitTime := caoGaoTime.Add(-minInterval)

	// 确保时间范围有效
	if !minSubmitTime.Before(maxSubmitTime) {
		// 如果时间太紧张，取中间点
		return payTime.Add(totalDuration / 2)
	}

	availableDuration := maxSubmitTime.Sub(minSubmitTime)
	submitOffset := time.Duration(rand.Int63n(int64(availableDuration)))
	submitTime := minSubmitTime.Add(submitOffset)

	return submitTime
}

func generatePublishTime(caoGaoTime time.Time) time.Time {
	rand.Seed(time.Now().UnixNano())

	// 最小间隔1天，最大间隔2天
	minOffset := 27 * time.Hour
	maxOffset := 48 * time.Hour

	// 获取 caoGaoTime 所在月份的最后一天
	lastDayOfMonth := time.Date(caoGaoTime.Year(), caoGaoTime.Month()+1, 0, 23, 59, 59, 0, caoGaoTime.Location())

	// 计算到月末的最大可用时间
	maxAvailableOffset := lastDayOfMonth.Sub(caoGaoTime)

	// 如果最大可用时间小于最小间隔，调整策略
	if maxAvailableOffset < minOffset {
		// 如果本月剩余时间不足，就设置为本月最后一天
		return lastDayOfMonth
	}

	// 调整最大偏移量，确保不跨月
	adjustedMaxOffset := maxOffset
	if maxAvailableOffset < maxOffset {
		adjustedMaxOffset = maxAvailableOffset
	}

	// 生成发布时间偏移
	publishOffset := minOffset + time.Duration(rand.Int63n(int64(adjustedMaxOffset-minOffset)))
	publishTime := caoGaoTime.Add(publishOffset)

	// 验证是否跨月（双重检查）
	if publishTime.Month() != caoGaoTime.Month() {
		// 如果跨月，就设置为本月最后一天
		publishTime = lastDayOfMonth
	}

	return publishTime
}

// 生成在指定时间之后的随机时间（随机分秒）
func generateRandomTimeAfter(baseTime time.Time) time.Time {
	// 生成随机分钟（1-120分钟）
	randomMinutes := time.Duration(1+rand.Intn(120)) * time.Minute

	// 生成随机秒数（0-59秒）
	randomSeconds := time.Duration(rand.Intn(60)) * time.Second

	totalOffset := randomMinutes + randomSeconds

	return baseTime.Add(totalOffset)
}
