package controllers

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func GetImage(ctx *gin.Context) {
	image_name := ctx.Param("image_name")
	ctx.File("./public/assets/images/" + image_name)
}

func UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" &&
			ext != ".gif" && ext != ".webp" {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{
					"status":  false,
					"message": "error",
					"data":    "file error",
				},
			)
			return
		} else {
			newFileName := time.Now().Format("") + "_" + file.Filename
			dst := filepath.Join("/public-v2/assets/images", newFileName)

			if err := ctx.SaveUploadedFile(file, dst); err == nil {
				// imageUrl := "path: /public/assets/images/<iamge_name>"
				imageUrl := dst
				ctx.AbortWithStatusJSON(
					http.StatusOK,
					gin.H{
						"status":  true,
						"message": "success",
						"data":    imageUrl,
					},
				)
				return
			} else {
				ctx.AbortWithStatusJSON(
					http.StatusBadRequest,
					gin.H{
						"status":  false,
						"message": "error",
						"data":    err.Error(),
					},
				)
				return
			}
		}
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  false,
				"message": "error",
				"data":    err.Error(),
			},
		)
		return
	}
}
