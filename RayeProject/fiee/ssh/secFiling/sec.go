package secFiling

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"test001/common"
	"test001/model"
	"time"

	"golang.org/x/crypto/ssh"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func PrSh() {
	// SSH 配置
	sshConfig := &ssh.ClientConfig{
		User: "liuzhihang",
		Auth: []ssh.AuthMethod{
			ssh.Password("liuzhihang001289391"), // 也可以用 ssh.PublicKeys(...)
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 测试用，生产要换安全方式
		Timeout:         time.Second * 5,
	}

	// 建立 SSH 连接
	sshClient, err := ssh.Dial("tcp", "121.40.49.103:22", sshConfig)
	if err != nil {
		panic("SSH连接失败: " + err.Error())
	}

	// 创建 MySQL 的远程连接转发
	mysqlHost := "hz001-szjixun.rwlb.rds.aliyuncs.com:3306" // 远程数据库地址
	localPort := "3307"                                     // 本地监听端口
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

	// 使用 GORM 连接本地转发端口（其实访问的是远程数据库）
	dsn := "fonchain_liuzhihang:JKAJKS1jdjjk182))@@tcp(127.0.0.1:3307)/micro-document?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	var failedRecord = make(map[int64]string)
	// 1. 读取文件
	content, err := os.ReadFile("secFiles.js")
	if err != nil {
		panic(err)
	}

	// 2. 去掉 JS 声明部分，保留 JSON 数组
	raw := string(content)
	raw = strings.TrimPrefix(raw, "export const fileList = ")
	raw = strings.TrimSuffix(raw, ";")

	// 3. 解析成 Go 对象
	var rawFilings []model.RawFiling
	if err := json.Unmarshal([]byte(raw), &rawFilings); err != nil {
		panic(err)
	}
	// 5. 插入数据
	for _, rf := range rawFilings {
		filingKey := common.GenerateFilingKey(rf.FilingDate, rf.Form, rf.FileLink)
		entry := model.SecFilings{
			FilingKey:       filingKey,
			FilingDate:      rf.FilingDate,
			Form:            rf.Form,
			Description:     rf.Description,
			FormDescription: rf.Description,
			FileLink:        rf.FileLink,
			DataFiles:       datatypes.JSON(rf.DataFiles),
			Status:          2,
		}

		// 避免重复插入（根据 FilingKey 查重）
		var count int64
		db.Model(&model.SecFilings{}).Where("filing_key = ?", filingKey).Count(&count)
		if count == 0 {
			if err := db.Create(&entry).Error; err != nil {
				fmt.Println("插入失败:", err)
				failedRecord[rf.Idx] = err.Error()
			} else {
				fmt.Println("插入成功:", filingKey)
			}
		} else {
			fmt.Println("跳过已存在记录:", filingKey)
			failedRecord[rf.Idx] = "filingKey已存在"
		}
	}
	// 保存路径
	filePath := "failed_records.json"

	// 转成 JSON 格式
	data, err := json.MarshalIndent(failedRecord, "", "  ")
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	// 写入文件
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	fmt.Println("✅ 失败记录已保存到:", filePath)
}

func copyConn(a, b net.Conn) {
	defer a.Close()
	defer b.Close()
	_, _ = io.Copy(a, b)
}
func FormType() {
	// SSH 配置
	sshConfig := &ssh.ClientConfig{
		User: "liuzhihang",
		Auth: []ssh.AuthMethod{
			ssh.Password("liuzhihang001289391"), // 也可以用 ssh.PublicKeys(...)
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 测试用，生产要换安全方式
		Timeout:         time.Second * 5,
	}

	// 建立 SSH 连接
	sshClient, err := ssh.Dial("tcp", "121.40.49.103:22", sshConfig)
	if err != nil {
		panic("SSH连接失败: " + err.Error())
	}

	// 创建 MySQL 的远程连接转发
	mysqlHost := "hz001-szjixun.rwlb.rds.aliyuncs.com:3306" // 远程数据库地址
	localPort := "3307"                                     // 本地监听端口
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

	// 使用 GORM 连接本地转发端口（其实访问的是远程数据库）
	dsn := "fonchain_liuzhihang:JKAJKS1jdjjk182))@@tcp(127.0.0.1:3307)/micro-document?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 自动迁移表结构（可选）
	//db.AutoMigrate(&model.FormType{})

	// 3. 读取现有数据库中所有 form_type 到 map[string]bool
	existingForms := make(map[string]bool)
	var allFormTypes []model.FormType
	if err := db.Model(&model.FormType{}).Select("form_type").Find(&allFormTypes).Error; err != nil {
		panic("查询数据库失败: " + err.Error())
	}
	for _, f := range allFormTypes {
		existingForms[f.FormType] = true
	}

	// 4. 打开 form.idx 文件
	file, err := os.Open("form.idx")
	if err != nil {
		panic("无法打开 form.idx 文件: " + err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seen := make(map[string]bool) // 本轮去重缓存
	var toInsert []model.FormType
	//now := time.Now().Unix()

	// 跳过前10行头部信息
	for i := 0; i < 10 && scanner.Scan(); i++ {
	}
	i := 0
	for scanner.Scan() {
		i++
		line := scanner.Text()
		// parts := strings.Fields(line)
		// if len(parts) == 0 {
		// 	continue
		// }
		form := strings.TrimRight(line[0:17], " ")
		// form := parts[0]
		if form == "" || seen[form] || existingForms[form] {
			continue
		}
		seen[form] = true
		toInsert = append(toInsert, model.FormType{
			FormType: form,
			//CreatedAt: now,
			//UpdatedAt: now,
		})
	}

	// 5. 插入新数据
	if len(toInsert) > 0 {
		if err := db.Create(&toInsert).Error; err != nil {
			panic("插入数据库失败: " + err.Error())
		}
		fmt.Printf("✅ 插入成功，共写入 %d 条新的 form_type。\n", len(toInsert))
	} else {
		fmt.Println("⚠️ 无需插入，所有 form_type 已存在。")
	}
	fmt.Println(i)
}
