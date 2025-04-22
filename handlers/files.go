package handlers

import (
	"net/http"
	"os"

	"backend/utils"

	"github.com/gin-gonic/gin"
)

func HandleFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не загружен"})
		return
	}

	// Создаем папку tmp, если её нет
	if err := os.MkdirAll("./tmp", 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания папки"})
		return
	}

	tempPath := "./tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения файла"})
		return
	}
	defer os.Remove(tempPath)

	// Отправляем файл в Telegram
	if err := utils.SendFileToTelegram(tempPath, "Файл от пользователя"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки файла"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Файл успешно отправлен!"})
}