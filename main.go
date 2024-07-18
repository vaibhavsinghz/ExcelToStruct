package main

import (
	"ExcelToStruct/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	router := gin.Default()

	router.POST("/upload", controller.UploadExcel)
}
