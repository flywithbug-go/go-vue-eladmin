package file_handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"vue-admin/web_server/model/model_file"

	"github.com/flywithbug/log4go"
)

func saveImageFile(file multipart.File, header *multipart.FileHeader) (imgPath string, err error) {
	defer file.Close()
	picture := model_file.Picture{}
	if statInterface, ok := file.(Size); ok {
		picture.Size = statInterface.Size()
	}
	if picture.Size > MaxPictureSize {
		err := fmt.Errorf("图片大小不能超过%s", MaxPictureSizeInfo)
		return "", err
	}
	//获取文件名
	ext := filepath.Ext(header.Filename)
	picture.Ext = ext

	//获取文件的md5值
	data, err := ioutil.ReadAll(file)
	picture.Md5 = makeMd5(data)
	fileName := picture.Md5 + ext

	//文件夹创建管理
	month := time.Now().Format("2006-01")
	localPath := localImageDirPath + month + "/"
	picture.Path = localPath

	localFilePath := localPath + fileName
	bExit, err := PathExists(localFilePath)

	if err != nil {
		return "", err
	}
	if bExit {
		avatarPath := fmt.Sprintf("/%s/%s", month, fileName)
		return avatarPath, nil
	}
	out, err := os.Create(localFilePath)
	if err != nil {
		bExit, err = PathExists(localPath)
		if err != nil {
			return "", err
		}
		if !bExit {
			err = os.Mkdir(localPath, os.ModePerm)
			if err != nil {
				err = os.Mkdir(localImageDirPath, os.ModePerm)
				if err != nil {
					return "", err
				}
				err = os.Mkdir(localPath, os.ModePerm)
				if err != nil {
					return "", err
				}
			}
		}
		//重新启动out
		out, err = os.Create(localFilePath)
		if err != nil {
			return "", err
		}
	}
	defer out.Close()
	_, err = out.Write(data)
	if err != nil {
		return "", err
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
	_, err = picture.Insert()
	if err != nil {
		log4go.Info(err.Error())
	}
	avatarPath := fmt.Sprintf("/%s/%s", month, fileName)
	return avatarPath, nil
}

func makeMd5(data []byte) string {
	h := md5.New()
	h.Write(data)
	value := h.Sum(nil)
	return hex.EncodeToString(value)
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
