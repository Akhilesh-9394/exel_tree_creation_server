package main

import (
	"gin_server/server/handler"
	"gin_server/server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	r.POST("/upload", handler.StudentHandler)
	r.GET("/:filename", handler.DownloadHandler)
	r.GET("/getstudents", handler.GetStudents)
	r.MaxMultipartMemory = 8 << 20
	r.Static("/assets", "./assets")
	r.Run(":3000")
}
