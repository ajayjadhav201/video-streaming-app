package controller

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	nfsBaseDir string = "/mnt/nfs/hls"
)

func VideoHandler(c *gin.Context) {
	hlsfilepath := c.Param("filepath")

	fullPath := filepath.Join(nfsBaseDir, hlsfilepath)

	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			c.String(http.StatusNotFound, "File not found")
			return
		}
		c.String(http.StatusInternalServerError, "Internal server error")
		return
	}
	if fileInfo.IsDir() {
		c.String(http.StatusForbidden, "Directory access is forbidden")
		return
	}

	file, err := os.Open(fullPath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer file.Close()

	c.Header("Content-Type", getContentType(&fullPath))
	c.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
	// c.Header("Accept-Ranges", "bytes").
	http.ServeContent(c.Writer, c.Request, fileInfo.Name(), fileInfo.ModTime(), file)

}

func getContentType(path *string) string {
	ext := strings.ToLower(filepath.Ext(*path))
	switch ext {
	case ".m3u8":
		return "application/vnd.apple.mpegurl"
	case ".ts":
		return "video/mp2t"
	default:
		return "application/octet-stream"
	}
}
