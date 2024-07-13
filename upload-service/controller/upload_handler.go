package controller

import (
	"path/filepath"
	"upload-service/utils"

	"github.com/gin-gonic/gin"
)

const (
	uploadDir = "/videos"
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		utils.WriteError(c, "Unable to retrive file", err)
		return
	}
	//
	startStr := c.PostForm("start")
	endStr := c.PostForm("end")
	//
	start := utils.Atoi(startStr, -1)
	if start == -1 {
		utils.WriteError(c, "Invalid start value", nil)
		return
	}
	end := utils.Atoi(endStr, -1)
	if end == -1 {
		utils.WriteError(c, "Invalid end value", nil)
		return
	}
	chunkFilename := utils.Sprintf("video_part_%d_%d%s",
		start, end, filepath.Ext(file.Filename),
	)

	chunkPath := filepath.Join(uploadDir, chunkFilename)

	if err := c.SaveUploadedFile(file, chunkPath); err != nil {
		utils.WriteError(c, "Failed to save file", err)
		return
	}
	utils.WriteJson(c, "Chunk uploaded successfully.")
}
