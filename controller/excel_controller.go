package controller

import (
	"ExcelToStruct/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadExcel(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}

	items, err := service.ConvertExcelToStruct(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}
