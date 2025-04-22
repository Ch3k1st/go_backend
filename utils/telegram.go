package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

func SendToTelegram(message string) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	client := resty.New()
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	_, err := client.R().
		SetFormData(map[string]string{
			"chat_id": chatID,
			"text":    message,
		}).
		Post(url)

	if err != nil {
		log.Printf("❌ Ошибка отправки в Telegram: %v", err)
		return err
	}

	log.Println("✅ Сообщение отправлено в Telegram")
	return nil
}

func SendFileToTelegram(filePath, caption string) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	client := resty.New()
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", botToken)

	_, err := client.R().
		SetFile("document", filePath).
		SetFormData(map[string]string{
			"chat_id": chatID,
			"caption": caption,
		}).
		Post(url)

	if err != nil {
		log.Printf("❌ Ошибка отправки файла: %v", err)
		return err
	}

	log.Println("✅ Файл отправлен в Telegram")
	return nil
}