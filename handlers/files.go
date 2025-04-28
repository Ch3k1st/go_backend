package handlers

import (
	"net/http"

	"backend/utils"

	"github.com/gin-gonic/gin"
)

func HandleFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не загружен"})
		return
	}

	if err := utils.SaveAndSendFile(c, file, "Файл от пользователя"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Файл успешно отправлен!"})
}