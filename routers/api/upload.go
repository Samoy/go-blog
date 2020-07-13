package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/pkg/app"
	"github.com/samoy/go-blog/pkg/e"
	"github.com/samoy/go-blog/pkg/logging"
	"github.com/samoy/go-blog/pkg/upload"
)

// UploadImage 上传图片
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.Error, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ErrorUploadCheckImageFormat, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ErrorUploadCheckImageFail, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ErrorUploadSaveImageFail, nil)
		return
	}

	appG.Response(http.StatusOK, e.Success, map[string]string{
		"image_url":      upload.GetImageFullURL(imageName),
		"image_save_url": savePath + imageName,
	})
}
