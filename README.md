# Sukhov-journal

Сервис статей, с возможностью их озвучки.

## Запуск проекта
```bash
git clone https://github.com/egorkagod/sukhov-journal.git
cd sukhov-journal
cp journal/.env.template journal/.env
docker compose up
```

## Tech Stack
- Go (Echo)
- Python (FastApi)
- PostgreSQL (Gorm)
- Docker / Docker Compose

## Архитектура

Проект состоит из двух сервисов, запущенных через Docker Compose:

1. **Articles Service (Go)** — основной бэкенд:
   - хранит статьи,
   - принимает запросы от клиента,
   - управляет пользователями и авторизацией,
   - отправляет текст на озвучивание.

2. **TTS/STT Service (Python)** — сервис обработки речи:
   - принимает текст от Go-сервиса,
   - генерирует аудио (Text-to-Speech),

Сервисы общаются между собой по HTTP внутри Docker-сети.