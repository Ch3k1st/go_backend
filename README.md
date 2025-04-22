# Telegram Form Backend

🚀 Простой сервер на Go, принимающий данные из формы и отправляющий их в Telegram.
Базовое и простое решение для получения информации из форм, к примеру, из реакт проекта.

## 📦 Стек

- Golang
- Gin
- Resty
- .env-переменные
- Telegram Bot API

## 🔧 Установка

1. Клонируйте репозиторий:

```bash
git clone https://github.com/Ch3k1st/go_backend
cd dir-repo
```

2. Установить зависимости:

```bash
go mod tidy
```

3. Создайте .env файл с зависимостями

```env
TELEGRAM_BOT_TOKEN=your_token_here
TELEGRAM_CHAT_ID=your_chat_id_here
```

4. Запуск сервера

Находясь в папке с проектом.

```bash
go run main.go
```

Сервер будет слушать порт :8080

📡 Эндпоинт
POST /send

Ожидаемые поля (в формате application/x-www-form-urlencoded):

- name
- surname
- direction
- email
- about

✅ Ответ

```json
{ "status": "Message sent" }
```

или

```json
{ "error": "Все поля должны быть заполнены!" }
```
