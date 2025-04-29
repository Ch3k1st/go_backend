package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

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

	c.JSON(http.StatusOK, gin.H{"status": "Message sent"})
}

// Форма для брендов
func HandleBrandForm(c *gin.Context) {
	brandName := c.PostForm("brandName")
	representative := c.PostForm("representative")
	phone := c.PostForm("phone")
	social := c.PostForm("social")
	email := c.PostForm("email")
	website := c.PostForm("website")
	productionType := c.PostForm("productionType")
	trademarkPatent := c.PostForm("trademarkPatent")

	// Валидация обязательных полей
	if brandName == "" || representative == "" || phone == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заполните обязательные поля!"})
		return
	}

	// Обработка файлов
	logoFile, _ := c.FormFile("logo")
	lookbookFile, _ := c.FormFile("lookbook")
	trademarkFile, _ := c.FormFile("trademarkFile")

	// Валидация форматов файлов (PNG/PDF)
	if logoFile != nil {
		ext := strings.ToLower(filepath.Ext(logoFile.Filename))
		if ext != ".png" && ext != ".pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Логотип должен быть PNG или PDF!"})
			return
		}
	}

	if lookbookFile != nil {
		ext := strings.ToLower(filepath.Ext(lookbookFile.Filename))
		if ext != ".png" && ext != ".pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Лукбук должен быть PNG или PDF!"})
			return
		}
	}

	if trademarkFile != nil {
		ext := strings.ToLower(filepath.Ext(trademarkFile.Filename))
		if ext != ".png" && ext != ".pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Свидетельство должно быть PNG или PDF!"})
			return
		}
	}

	message := fmt.Sprintf(`
	🏢 *Новая заявка от Бренда* 🏢
	———————————————
	🔹 *Название бренда:* %s
	🔹 *Представитель:* %s
	🔹 *Телефон:* %s
	🔹 *Соцсети:* %s
	🔹 *Email:* %s
	🔹 *Сайт:* %s
	🔹 *Производство:* %s
	🔹 *Патент на товарный знак:* %s
	`, brandName, representative, phone, social, email, website, productionType, trademarkPatent)

	if err := utils.SendToTelegram(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки в Telegram"})
		return
	}

	// Отправка файлов
	if logoFile != nil {
		if err := utils.SaveAndSendFile(c, logoFile, "Логотип бренда"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки логотипа: " + err.Error()})
			return
		}
	}

	if lookbookFile != nil {
		if err := utils.SaveAndSendFile(c, lookbookFile, "Лукбук коллекции"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки лукбука: " + err.Error()})
			return
		}
	}

	if trademarkFile != nil {
		if err := utils.SaveAndSendFile(c, trademarkFile, "Свидетельство на товарный знак"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки свидетельства: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "✅ Форма бренда успешно отправлена!"})
}

// Форма для дизайнеров
func HandleDesignerForm(c *gin.Context) {
	fullName := c.PostForm("fullName")
	birthDate := c.PostForm("birthDate")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	social := c.PostForm("social")
	experience := c.PostForm("experience")
	awards := c.PostForm("awards")

	// Валидация обязательных полей
	if fullName == "" || phone == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ФИО, телефон и email обязательны!"})
		return
	}

	// Обработка файлов
	logoFile, _ := c.FormFile("logo")
	lookbookFile, _ := c.FormFile("lookbook")

	// Проверка форматов файлов (PNG/PDF)
	if logoFile != nil {
		ext := strings.ToLower(filepath.Ext(logoFile.Filename))
		if ext != ".png" && ext != ".pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Логотип должен быть PNG или PDF!"})
			return
		}
	}

	if lookbookFile != nil {
		ext := strings.ToLower(filepath.Ext(lookbookFile.Filename))
		if ext != ".png" && ext != ".pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Портфолио должно быть PNG или PDF!"})
			return
		}
	}

	message := fmt.Sprintf(`
	🎨 *Новая заявка от Дизайнера* 🎨
	———————————————
	🔹 *ФИО:* %s
	🔹 *Дата рождения:* %s
	🔹 *Телефон:* %s
	🔹 *Email:* %s
	🔹 *Соцсети:* %s
	🔹 *Опыт работы:* %s
	🔹 *Награды:* %s
	`, fullName, birthDate, phone, email, social, experience, awards)

	if err := utils.SendToTelegram(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки текста в Telegram: " + err.Error()})
		return
	}

	if logoFile != nil {
		if err := utils.SaveAndSendFile(c, logoFile, "Логотип дизайнера"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки логотипа: " + err.Error()})
			return
		}
	}

	if lookbookFile != nil {
		if err := utils.SaveAndSendFile(c, lookbookFile, "Портфолио дизайнера"); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка отправки портфолио: " + err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "✅ Форма дизайнера успешно отправлена!"})
}