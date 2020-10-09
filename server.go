package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	InitDB()
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	data := DataHandler{}
	r.GET("/ping", Ping)
	r.POST("/DataCleansing/", data.EndPointDataCleansing)
	return r
}