package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samoy/go-blog/pkg/e"
	"github.com/samoy/go-blog/pkg/logging"
	"github.com/samoy/go-blog/pkg/upload"
)

// UploadImage 上传图片
func UploadImage(c *gin.Context) {
	code := e.Success
	data := make(map[string]string)
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.Error
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.InvalidParams
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()
		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ErrorUploadCheckImageFormat
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ErrorUploadCheckImageFail
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ErrorUploadSaveImageFail
			} else {
				data["image_url"] = upload.GetImageFullURL(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}
}
