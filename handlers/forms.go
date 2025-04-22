package handlers

import (
	"fmt"
	"net/http"

	"backend/utils"

	"github.com/gin-gonic/gin"
)

func HandleFormSubmit(c *gin.Context) {
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	direction := c.PostForm("direction")
	email := c.PostForm("email")
	about := c.PostForm("about")

	if name == "" || surname == "" || direction == "" || email == "" || about == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Все поля должны быть заполнены!"})
		return
	}

	message := fmt.Sprintf(`
	📝 Новая заявка!
	──────────────
	ФИО: %s %s
	Направление: %s
	Email: %s
	О себе: %s
	`, name, surname, direction, email, about)

	if err := utils.SendToTelegram(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки в Telegram"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Успешно отправлено!"})
}