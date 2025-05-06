package logic

import (
	"chain-dci/pb/dci"
	"chain-dci/pkg/app"
	errCommon "chain-dci/pkg/err"
	"chain-dci/pkg/msg"
	"chain-dci/pkg/utils"
	"encoding/base64"
	"errors"
	"fmt"
	bccrClient "github.com/antchain-openapi-sdk-go/bccr/client"
	"github.com/jinzhu/copier"
	"strings"
	"time"
	"unicode"
)

type IFile interface {
	Upload(request *dci.GetUploadUrlRequest) error
}

type File struct {
}

/*
		存证API
		GetUploadUrl
	    获取 访问 oss的链接, 注意 如果文件名为中文 则需要 对文件名使用 utf-8 字符集进行 URL编码上传
*/
func GetUploadUrl(req *bccrClient.GetUploadurlRequest) (result *bccrClient.GetUploadurlResponse) {

	result, err := app.ModuleClients.BccrClient.GetUploadurl(req)
	if err != nil {
		errCommon.NoReturnError(err, "调用 蚂蚁链获取 oss url 错误:")
		return
	}
	errCommon.NoReturnInfo(req, "获取 DCI GetUploadurl 返回结果:")
	return
}

/*
		存证API
		Upload
	    上传 实际文件
*/
func (f *File) Upload(req *dci.GetUploadUrlRequest) (res *dci.GetUploadUrlResponse, err error) {

	fmt.Println("+++++++++++++++ GetUploadUrlRequest  =================")
	fmt.Printf("GetUploadUrlRequest is : %+v\n", req)
	fmt.Println("+++++++++++++++ GetUploadUrlRequest  =================")

	errCommon.NoReturnInfo(req, "上传文件 参数记录:")

	/*// 记录 上传的文件
	//nil := app.ModuleClients.DciDB.Begin()
	fileInfo := new(model.FileInfo)
	fileInfo.FileUrl = req.FileName
	_ = dao.CreateFileInfo(nil, fileInfo)*/

	isChinese := false

	getUploadUrlRequest := new(bccrClient.GetUploadurlRequest)

	var fileName string
	// 拆分 文件名 和 文件链接
	fileArr := strings.Split(req.FileName, "/")
	fileName = fileArr[len(fileArr)-1]
	fmt.Println("+++++++++++++++ no encode  fileName  =================")
	fmt.Println("no encode  fileName is :", fileArr[len(fileArr)-1])
	fmt.Println("+++++++++++++++ no encode  fileName  =================")

	//  如果文件名包含 中文 则需要对 fileName 进行 utf-8 字符集进行 URL编码
	for _, c := range fileName {
		if unicode.Is(unicode.Scripts["Han"], c) {
			isChinese = true
		}
	}

	if isChinese {
		fileName = base64.URLEncoding.EncodeToString([]byte(fileName))
	}
	getUploadUrlRequest.SetFileName(fileName)

	clientToken, err := createToken(time.Now().UnixMilli(), fileName, app.ModuleClients.SfNode.Generate().Base64())
	if err != nil {
		return nil, errCommon.ReturnError(err, msg.ErrCreateClientToken, "创建clientToken 失败: ")
	}
	getUploadUrlRequest.SetClientToken(clientToken)
	fmt.Println("+++++++++++++++ encode  fileName  =================")
	fmt.Println("encode  fileName is :", fileName)
	fmt.Println("+++++++++++++++ encode  fileName  =================")

	getUploadUrlResponse := GetUploadUrl(getUploadUrlRequest)
	res = new(dci.GetUploadUrlResponse)
	_ = copier.CopyWithOption(&res, getUploadUrlResponse, copier.Option{DeepCopy: false})

	fmt.Println("+++++++++++++++ GetUploadUrl  =================")
	fmt.Printf("GetUploadUrl is : %+v\n", res)
	fmt.Println("+++++++++++++++ GetUploadUrl  =================")

	fmt.Println("====== =========== ===================     1   ===============")
	if res.ResultCode != "OK" {
		/*	fileInfo.ReqMsgId = res.ReqMsgId
			fileInfo.ResultCode = res.ResultCode
			fileInfo.ResultMsg = res.ResultMsg
			fileInfo.Url = res.Url
			fileInfo.FileId = res.FileId
			err = dao.UpdateFileInfo(nil, fileInfo)*/
		errCommon.NoReturnError(errors.New(res.ResultMsg), "获取授权访问OSS链接 错误:")
		return res, nil
	}

	fmt.Println("====== =========== ===================     2   ===============")
	if res.Url == "" || res.FileId == "" {
		/*fileInfo.ReqMsgId = res.ReqMsgId
		fileInfo.ResultCode = res.ResultCode
		fileInfo.ResultMsg = res.ResultMsg
		fileInfo.Url = res.Url
		fileInfo.FileId = res.FileId
		_ = dao.UpdateFileInfo(nil, fileInfo)*/
		errCommon.NoReturnError(errors.New(res.ResultMsg), "获取授权访问OSS链接 错误:")
		return res, nil
	}

	fmt.Println("====== =========== ===================     3   ===============")
	code, result := utils.PutFromFileUrlWithStream(res.Url, fileName, req.FileName)
	if code != 200 {
		err = errCommon.ReturnError(errors.New(result), result, "上传文件 错误:")
		//fileInfo.IsUpload = 3 // 上传失败
	} else {
		errCommon.NoReturnInfo(result, "上传文件 最终结果:")
		//fileInfo.IsUpload = 2 // 上传成功
	}
	/*fileInfo.ReqMsgId = res.ReqMsgId
	fileInfo.ResultCode = res.ResultCode
	fileInfo.ResultMsg = res.ResultMsg
	fileInfo.Url = res.Url
	fileInfo.FileId = res.FileId
	_ = dao.UpdateFileInfo(nil, fileInfo)*/

	return res, err
}
