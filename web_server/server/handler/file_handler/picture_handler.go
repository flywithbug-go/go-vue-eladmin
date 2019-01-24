package file_handler

import (
	"errors"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"vue-admin/web_server/common"
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

const (
	MaxPictureSize     int64 = 10485760
	MaxPictureSizeInfo       = "10m"
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

	imgPath, err := saveImageFile(file, header)
	if err != nil {
		log4go.Info(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusBadRequest, fmt.Sprintf("write file err : %s", err.Error()))
		return
	}
	aRes.SetResponseDataInfo("imagePath", imgPath)
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

func scale(in io.Reader, out io.Writer, size, quality int) error {
	origin, fm, err := image.Decode(in)
	if err != nil {
		return err
	}
	width := size
	height := 0
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
