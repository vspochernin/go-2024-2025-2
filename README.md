# Банковская система

REST API для банковского сервиса на Go.

## Требования

- Go 1.23+
- Docker и Docker Compose
- PostgreSQL 17

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
```

2. Запустите базу данных:
```bash
docker-compose up -d
```

3. Установите зависимости:
```bash
go mod download
```

4. Настройте переменные окружения:
```bash
cp .env.example .env
# Отредактируйте .env файл под свои нужды
```

5. Запустите приложение:
```bash
go run cmd/api/main.go
```

## Структура проекта

```
.
├── cmd/
│   └── api/           # Точка входа приложения
├── internal/
│   ├── models/        # Модели данных
│   ├── repositories/  # Работа с БД
│   ├── services/      # Бизнес-логика
│   ├── handlers/      # Обработчики HTTP
│   └── middleware/    # Middleware
├── migrations/        # SQL миграции
├── config/           # Конфигурация
└── docker-compose.yml # Конфигурация Docker
```

## API Endpoints

### Публичные
- POST /register - Регистрация пользователя
- POST /login - Аутентификация

### Защищенные
- POST /accounts - Создание счета
- POST /cards - Выпуск карты
- POST /transfer - Перевод средств
- GET /analytics - Получение аналитики
- GET /credits/{creditId}/schedule - График платежей
- GET /accounts/{accountId}/predict - Прогноз баланса

## Технологии

- Go 1.23
- PostgreSQL
- JWT для аутентификации
- PGP для шифрования
- SMTP для уведомлений
- SOAP для интеграции с ЦБ РФ
