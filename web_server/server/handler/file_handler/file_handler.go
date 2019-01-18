package file_handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model/model_file"
	"vue-admin/web_server/server/handler/handler_common"

	"github.com/flywithbug/log4go"

	"vue-admin/web_server/model"

	"github.com/gin-gonic/gin"
)

var (
	localFileDirPath = "../file/"
)

func SetLocalFilePath(path string) {
	localFileDirPath = path
}

// 获取文件大小的接口
type Size interface {
	Size() int64
}

// 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

func uploadFileHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponseCode, aRes.Code)
		c.JSON(http.StatusOK, aRes)
	}()

	//gin将het/http包的FormFile函数封装到c.Request
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, fmt.Sprintf("get file err : %s", err.Error()))
		return
	}
	defer file.Close()

	localFile := model_file.File{}
	if statInterface, ok := file.(Size); ok {
		localFile.Size = statInterface.Size()
	}
	//获取文件名
	ext := filepath.Ext(header.Filename)
	localFile.Ext = ext
	//获取文件的md5值
	data, err := ioutil.ReadAll(file)
	h := md5.New()
	h.Write(data)
	value := h.Sum(nil)
	localFile.Md5 = hex.EncodeToString(value)
	fileName := localFile.Md5 + ext

	//文件夹创建管理
	month := time.Now().Format("2006-01")
	localPath := localFileDirPath + month + "/"
	localFile.Path = localPath

	//文件路径
	localFilePath := localPath + fileName
	bExit, err := PathExists(localFilePath)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, fmt.Sprintf("system err : %s", err.Error()))
		return
	}
	if bExit {
		log4go.Info(handler_common.RequestId(c)+"fileExit: %s", fileName)
		avatarPath := fmt.Sprintf("/%s/%s", month, fileName)
		aRes.SetResponseDataInfo("imagePath", avatarPath)
		return
	}
	out, err := os.Create(localFilePath)
	if err != nil {
		log4go.Info(handler_common.RequestId(c)+"创建文件失败：%s", err.Error())
		//判断文件夹是否存在
		bExit, err = PathExists(localPath)
		if err != nil {
			log4go.Info(handler_common.RequestId(c) + err.Error())
			aRes.SetErrorInfo(http.StatusInternalServerError, fmt.Sprintf("get folder err : %s", err.Error()))
			return
		}
		//文件夹不存在创建文件夹
		if !bExit {
			err = os.Mkdir(localPath, os.ModePerm)
			if err != nil {
				log4go.Info(handler_common.RequestId(c) + err.Error())
				aRes.SetErrorInfo(http.StatusInternalServerError, fmt.Sprintf("make folder err : %s", err.Error()))
				return
			}
		}
		//重新启动out
		out, err = os.Create(localFilePath)
		if err != nil {
			log4go.Info(handler_common.RequestId(c) + err.Error())
			aRes.SetErrorInfo(http.StatusInternalServerError, fmt.Sprintf("make file err : %s", err.Error()))
			return
		}
	}
	defer out.Close()

	_, err = out.Write(data)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, fmt.Sprintf("write file err : %s", err.Error()))
		return
	}
	localFile.Insert()
	filePath := fmt.Sprintf("/%s/%s", month, fileName)
	aRes.SetResponseDataInfo("filePath", filePath)
}
