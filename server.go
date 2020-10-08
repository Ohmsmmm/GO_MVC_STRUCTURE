package main

import (
	"github.com/GO_MVC_STRUCTURE/controller"
	"github.com/gin-gonic/gin"
)


func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	data := controller.DataHandler{}
	r.GET("/ping", controller.Ping)
	r.POST("/DataCleansing/:input_string", data.EndPointDataCleansing)
	return r
}