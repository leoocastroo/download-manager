package main

import (
	"github.com/gin-gonic/gin"
	"download-manager/controller"
	"download-manager/config"
	"download-manager/model"
)


func main() {
	config.LoadEnv() // load the enviroment

	r := gin.Default()

	model.ConnectDatabase()

	r.GET("/get-files-info", controller.GetFilesInfo)
	r.POST("/download-file", controller.DownloadFile)
	r.POST("/download-files", controller.DownloadFiles)
	r.Run()
}
