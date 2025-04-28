package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Роуты
	r.POST("/send-form", handlers.HandleFormSubmit)            // Форма под маленький joinUs
	r.POST("/send-brand-form", handlers.HandleBrandForm)       // Форма бренда
	r.POST("/send-designer-form", handlers.HandleDesignerForm) // Форма дизайнера

	log.Println("✅ Сервер запущен на порту 8080")
	r.Run(":8080")
}
