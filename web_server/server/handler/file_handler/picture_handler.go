package file_handler

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model/model_file"
	"vue-admin/web_server/server/handler/handler_common"

	"github.com/flywithbug/file"
	"github.com/flywithbug/log4go"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"

	"vue-admin/web_server/model"

	"github.com/gin-gonic/gin"

	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var (
	localImageDirPath = "../image/"
)

func SetLocalImageFilePath(path string) {
	localImageDirPath = path
}

func uploadImageHandler(c *gin.Context) {
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
	picture := model_file.Picture{}
	if statInterface, ok := file.(Size); ok {
		picture.Size = statInterface.Size()
	}
	//10M
	if picture.Size > 10485760 {
		log4go.Info(handler_common.RequestId(c) + "picture size > 10m")
		aRes.SetErrorInfo(http.StatusBadRequest, "picture size > 10m")
		return
	}
	//获取文件名
	ext := filepath.Ext(header.Filename)
	picture.Ext = ext
	//获取文件的md5值
	data, err := ioutil.ReadAll(file)

	h := md5.New()
	h.Write(data)
	value := h.Sum(nil)
	picture.Md5 = hex.EncodeToString(value)
	fileName := picture.Md5 + ext

	//文件夹创建管理
	month := time.Now().Format("2006-01")
	localPath := localImageDirPath + month + "/"
	picture.Path = localPath

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
	pictureFile, err := os.Open(localFilePath)
	if err != nil {
		log4go.Info(err.Error())
	}
	imgConf, _, err := image.DecodeConfig(pictureFile)
	if err != nil {
		log4go.Info(err.Error())
	}
	picture.Width = imgConf.Width
	picture.Height = imgConf.Height
	picture.Insert()
	avatarPath := fmt.Sprintf("/%s/%s", month, fileName)
	aRes.SetResponseDataInfo("imagePath", avatarPath)
}

func loadImageHandler(c *gin.Context) {
	path := c.Param("path")
	filename := c.Param("filename")
	//log4go.Info(handler_common.RequestId(c) + "loadImageHandler: %s %s", path, filename)
	if path == "" || filename == "" {
		return
	}
	size := c.Query("size")

	fileOrigin := localImageDirPath + path + "/" + filename
	sizeW, err := strconv.Atoi(size)

	if len(size) == 0 || err != nil {
		http.ServeFile(c.Writer, c.Request, fileOrigin)
		return
	}
	ext := filepath.Ext(filename)
	if strings.EqualFold(ext, ".gif") {
		http.ServeFile(c.Writer, c.Request, fileOrigin)
		return
	}

	filePath := localImageDirPath + path + "/" + size + "-" + filename
	if !file.FileExists(filePath) {
		if !file.FileExists(fileOrigin) {
			c.Writer.Write([]byte("Error: Image Not found."))
			return
		}
		fIn, _ := os.Open(fileOrigin)
		//log4go.Info(handler_common.RequestId(c) + fileOrigin)
		defer fIn.Close()
		fOut, _ := os.Create(filePath)
		//log4go.Info(handler_common.RequestId(c) + filename)
		defer fOut.Close()
		if sizeW < 100 {
			sizeW = 100
		}
		err := scale(fIn, fOut, sizeW, 100)
		if err != nil {
			log4go.Info(handler_common.RequestId(c) + err.Error())
			http.ServeFile(c.Writer, c.Request, fileOrigin)
			return
		}
	}
	http.ServeFile(c.Writer, c.Request, filePath)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func scale(in io.Reader, out io.Writer, size, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		return err
	}
	width := size
	height := size
	if size == 0 {
		width = origin.Bounds().Max.X
		height = origin.Bounds().Max.Y
	} else {
		height = origin.Bounds().Max.Y * (size / origin.Bounds().Max.X)
	}

	if quality == 0 {
		quality = 100
	}
	canvas := resize.Thumbnail(uint(width), uint(height), origin, resize.Lanczos3)

	switch fm {
	case "jpeg":
		return jpeg.Encode(out, canvas, &jpeg.Options{Quality: quality})
	case "png":
		return png.Encode(out, canvas)
	case "gif":
		return gif.Encode(out, canvas, &gif.Options{})
	case "bmp":
		return bmp.Encode(out, canvas)
	default:
		return errors.New("ERROR FORMAT")
	}
	return nil
}
