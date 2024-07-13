package main

import (
	"watch-service/utils"

	"github.com/gin-gonic/gin"
)

var (
	HostUrl = ":8082"
)

func main() {
	//
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	//
	utils.FatalIfError(r.Run(HostUrl), "")
}
