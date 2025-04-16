# Разработка банковской системы - задание

## Введение

Проект представляет собой разработку REST API для банковского сервиса на языке Go. Цель - создать удобное и безопасное приложение для доступа к банковским услугам.

В приложении будет интеграция с внешними сервисами - Центральным банком РФ и SMTP. Особое внимание будет уделено безопасности: шифрованию данных банковских карт и использованию методов аутентификации, таких как JWT.

Приложение должно состоять из нескольких слоев:

- Модели данных.
- Репозитории для работы с базой данных.
- Сервисы для реализации бизнес-логики.
- Обработчики запросов.

Каждый из слоев выполняет свою роль, обеспечивая чистую архитектуру и легкость в сопровождении кода.

Данный проект закрепляет знания по созданию серверных приложений на Go, а также предоставляет практику в ключевых темах: создании REST API, взаимодействии с базой данных, настройке аутентификации и защите информации.

## Что нужно сделать

### Функциональные требования к REST API банковского сервиса

В рамках проекта необходимо реализовать REST API для банковского сервиса со следующими функциями:

- Регистрация пользователей с проверкой уникальности.
- Аутентификация пользователей.
- Создание банковских счетов и управление ими.
- Операции с картами: генерация, просмотр, оплата
- Переводы между счетами и пополнение баланса.
- Кредитные операции: оформление кредита, график платежей.
- Аналитика финансовых операций.
- Интеграция с внешними сервисами:
- Центральный банк РФ — для определения ключевой ставки;
- SMTP — для отправки уведомлений по электронной почте.

Также должна быть реализована система защиты данных, включая шифрование и хеширование.

### Рекомендации по использованию библиотек

- Язык: Go 1.23+.
- Фреймворки/библиотеки:
- Маршрутизация: gorilla/mux.
- Работа с БД: PostgreSQL + lib/pq.
- Аутентификация: JWT (golang-jwt/jwt/v5).
- Логирование: logrus.
- Шифрование: bcrypt, HMAC-SHA256, PGP.
- Дополнительно: gomail.v2 (отправка email), beevik/etree (парсинг XML).
- База данных: PostgreSQL 17 с расширением pgcrypto.

### Требования к структуре приложения

Приложение должно состоять из следующих основных слоев.

#### Модели

- Определение структур данных (соответствие таблицам БД).
- Валидация полей (email, пароль, уникальность).
- Сериализация/десериализация для API (теги JSON).

#### Репозитории

- Инкапсуляция SQL-запросов.
- Параметризованные запросы (защита от инъекций).
- Обработка ошибок БД, управление транзакциями.

#### Сервисы

- Реализация бизнес-логики (переводы, кредиты, аналитика).
- Интеграция с внешними сервисами (SMTP, API ЦБ РФ).
- Обработка ошибок, возврат понятных сообщений.

#### Обработчики (Handlers)

- Валидация входных данных запросов.
- Вызов методов сервисов.
- Формирование HTTP-ответов (статусы, JSON).
- Использование middleware для безопасности.

#### Маршрутизация

- Определение эндпоинтов (методы, пути).
- Разделение публичных и защищенных маршрутов.
- Подключение middleware (аутентификация).

#### Middleware

- Проверка JWT-токенов.
- Добавление контекста (например, ID пользователя).
- Блокировка неавторизованных запросов.

### Функциональные требования

#### Пользовательские операции

- Регистрация (с проверкой уникальности email и username).
- Аутентификация (JWT с сроком действия 24 часа).

#### Операции с банковскими счетами

- Создание банковских счетов.
- Пополнение баланса / Списание средств со счета.
- Переводы между счетами.

#### Операции с картами

- Генерация виртуальных карт (с валидным номером по алгоритму Луна).
- Хранение данных карт в зашифрованном виде (PGP) и HMAC.
- Просмотр карт (с расшифровкой для владельца).

#### Кредитные операции

- Оформление кредита с расчетом аннуитетных платежей.
- Автоматическое списание платежей (шедулер каждые N часов).
- Генерация графика платежей.
- Начисление штрафов за просрочку, если на счету кредита нет достаточно средств (+10% к сумме).

#### Предоставление аналитики для клиента

- Статистика по доходам/расходам за месяц.
- Аналитика кредитной нагрузки.
- Прогноз баланса на N дней (учет запланированных платежей).

#### Интеграции

- Получение ключевой ставки ЦБ РФ через SOAP-запрос.
- Отправка email-уведомлений о платежах (SMTP).

### Примерная структура эндпоинтов приложения

#### Публичные

- POST /register — регистрация.
- POST /login — аутентификация.

#### Защищенные (требуют JWT реализованного в middleware)

- POST /accounts — создать счет.
- POST /cards — выпустить карту.
- POST /transfer — перевод средств.
- GET /analytics — получить аналитику.
- GET /credits/{creditId}/schedule — график платежей по кредиту.
- GET /accounts/{accountId}/predict — прогноз баланса.

И другие эндпоинты

### Требования к базе данных

Как минимум, должны присутствовать следующие таблицы:

- users: данные пользователей.
- accounts: банковские счета.
- cards: данные карт (номер, срок, CVV в зашифрованном виде).
- transactions: история операций.
- credits: кредиты.
- payment_schedules: график платежей.

### Требования к безопасности

- Номер и срок карты шифруются PGP с ключом.
- CVV хешируется через bcrypt.
- HMAC используется для проверки целостности данных.
- Аутентификация: JWT с секретом (переменная JWT_SECRET).
- Данные карт: шифрование PGP + HMAC.
- Пароли: хеширование bcrypt.
- Проверка прав доступа к счетам/картам (например, запрет пополнения чужого счета).

### Дополнительные требования

- Шедулер: обработка просроченных платежей каждые 12 часов.
- Интеграция с ЦБ РФ: SOAP-запросы к https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx для получения актуальной ключевой ставки.
- Логирование: через logrus (уровень логирования можно настроить).

### Ограничения

- Поддержка только RUB в качестве валюты.
- Максимальный период прогноза баланса: 365 дней.

## Методические инструкции

### Как настроить интеграцию с ЦБ РФ

Для удобной работы с XML-ответами используется библиотека etree. Ее можно установить следующим образом:

```shell
go get github.com/beevik/etree
```

Далее нужно добавить необходимые импорты в Go-файл:

```go
import (
    "bytes"
    "encoding/xml"
    "fmt"
    "io"
    "net/http"
    "time"
    "github.com/beevik/etree"
)
```

ЦБ РФ предоставляет SOAP-API. Пример формирования запроса:

```go
func buildSOAPRequest() string {
    fromDate := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
    toDate := time.Now().Format("2006-01-02")
    return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
        <soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
            <soap12:Body>
                <KeyRate xmlns="http://web.cbr.ru/">
                    <fromDate>%s</fromDate>
                    <ToDate>%s</ToDate>
                </KeyRate>
            </soap12:Body>
        </soap12:Envelope>`, fromDate, toDate)
}
```

Нужно настроить HTTP-клиент и отправить запрос:

```go
func sendRequest(soapRequest string) ([]byte, error) {
    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest(
        "POST",
        "https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx",
        bytes.NewBuffer([]byte(soapRequest)),
    )
    if err != nil {
        return nil, err
    }
    // Установка заголовков
    req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
    req.Header.Set("SOAPAction", "http://web.cbr.ru/KeyRate")

    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("ошибка запроса: %v", err)
    }
    defer resp.Body.Close()
    // Чтение ответа
    rawBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("ошибка чтения ответа: %v", err)
    }

    return rawBody, nil
}
```

Для парсинга полученного ответа можно использовать etree:

```go
func parseXMLResponse(rawBody []byte) (float64, error) {
    doc := etree.NewDocument()
    if err := doc.ReadFromBytes(rawBody); err != nil {
        return 0, fmt.Errorf("ошибка парсинга XML: %v", err)
    }
    // Поиск элементов в XML
    krElements := doc.FindElements("//diffgram/KeyRate/KR")
    if len(krElements) == 0 {
        return 0, errors.New("данные по ставке не найдены")
    }
    latestKR := krElements[0]
    rateElement := latestKR.FindElement("./Rate")
    if rateElement == nil {
        return 0, errors.New("тег Rate отсутствует")
    }
    // Конвертация строки в число
    rateStr := rateElement.Text()
    var rate float64
    if _, err := fmt.Sscanf(rateStr, "%f", &rate); err != nil {
        return 0, fmt.Errorf("ошибка конвертации ставки: %v", err)
    }

    return rate, nil
}
```

Наконец, нужно добавить маржу банка к ставке (например, +5%):

```go
func GetCentralBankRate() (float64, error) {
    soapRequest := buildSOAPRequest()
    rawBody, err := sendRequest(soapRequest)
    if err != nil {
        return 0, err
    }
    rate, err := parseXMLResponse(rawBody)
    if err != nil {
        return 0, err
    }

    // Добавление маржи
    rate += 5
    return rate, nil
}
```

### Как настроить интеграцию с SMTP

SMTP (Simple Mail Transfer Protocol) — стандартный протокол для отправки электронной почты. Интеграция с SMTP-сервером позволяет приложению автоматизированно отправлять уведомления пользователям по электронной почте.

Для работы с SMTP в Go рекомендуется использовать библиотеку gomail:

```shell
go get github.com/go-mail/mail/v2
```

Нужно импортировать необходимые зависимости:

```go
import (
    "crypto/tls"
    "fmt"
    "log"

    "github.com/go-mail/mail/v2"
)
```

Настроить параметры подключения:

```go
const (
    smtpHost = "smtp.example.com"  // Хост SMTP-сервера
    smtpPort = 587                 // Порт (чаще используется 587 с TLS)
    smtpUser = "noreply@example.com" // Учетная запись
    smtpPass = "strong_password"    // Пароль/токен
)
```

Можно использовать любой почтовый сервис для рассылки писем, например mailgun.

Далее нужно сформировать сообщение для отправки по почте:

```go
func createMessage(to string, subject string, body string) *mail.Message {
    m := mail.NewMessage()
    m.SetHeader("From", smtpUser)
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)
    return m
}
```

После этого настроить SMTP-диалог:

```go
func createDialer() *mail.Dialer {
    d := mail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
    d.TLSConfig = &tls.Config{
        ServerName:         smtpHost,
        InsecureSkipVerify: false, // Не отключать проверку сертификата
    }
    return d
}
```

Отправить письмо:

```go
func sendEmail(d *mail.Dialer, m *mail.Message) error {
    if err := d.DialAndSend(m); err != nil {
        log.Printf("SMTP error: %v", err)
        return fmt.Errorf("email sending failed")
    }
    return nil
}
```

Итоговая реализация выглядит следующим образом:

```go
func (s *BankService) sendPaymentEmail(userEmail string, amount float64) error {
    // Создание контента
    content := fmt.Sprintf(`
        <h1>Спасибо за оплату!</h1>
        <p>Сумма: <strong>%.2f RUB</strong></p>
        <small>Это автоматическое уведомление</small>
    `, amount)
    // Подготовка сообщения
    m := createMessage(userEmail, "Платеж успешно проведен", content)

    // Настройка подключения
    d := createDialer()

    // Отправка
    if err := sendEmail(d, m); err != nil {
        return err
    }

    log.Printf("Email sent to %s", userEmail)
    return nil
}
```

### Как настроить JWT-аутентификацию

Сначала нужно установить необходимую библиотеку:

```shell
go get github.com/golang-jwt/jwt/v5
```

Импортировать ее в проект:

```go
import (
	"context"
	"net/http"
	"strings"
"github.com/golang-jwt/jwt/v5"
)
```

Настроить процесс проверки JWT при получении каждого запроса, требующего авторизации:

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header required", http.StatusUnauthorized)
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        claims := &jwt.RegisteredClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims,
            func(token *jwt.Token) (interface{}, error) {
                return []byte(JWT_SECRET), nil
            })
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), "userID", claims.Subject)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

Теперь нужно настроить процесс генерации JWT (например, когда пользователь логинится):

```go
func GenerateJWTToken(userID string) (string, error) {
    claims := jwt.RegisteredClaims{
        Subject:   userID,
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(JWT_SECRET))
}
```

Осталось осуществить разграничение маршрутов эндпоинтов на требующие и не требующие авторизации:

```go
// Предшествующий код
// Public routes
r.HandleFunc("/register", h.Register).Methods("POST")
r.HandleFunc("/login", h.Login).Methods("POST")
// Protected routes
authRouter := r.PathPrefix("/").Subrouter()
authRouter.Use(AuthMiddleware)

// Управление счетами
authRouter.HandleFunc("/accounts", h.CreateAccount).Methods("POST")
authRouter.HandleFunc("/accounts", h.GetUserAccounts).Methods("GET")
// Дальнейший код
```

### Использование библиотек шифрования

Для шифрования паролей можно использовать библиотеку crypto:

```shell
go get golang.org/x/crypto
```

Импортировать пакет bcrypt:

```go
import (
	"golang.org/x/crypto/bcrypt"
)
```

Использовать ее для реализации следующих операций:

- Хеширование пароля при регистрации.

```go
hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```

- Проверка пароля при логине.


```go
err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
```

- Хеширование CVV карты.

```go
func HashCVV(cvv string) (string, error) {
    return bcrypt.GenerateFromPassword([]byte(cvv), bcrypt.DefaultCost)
}
```

Для защиты данных банковских карт и их подтверждения стоит использовать HMAC-SHA256, подключив пакеты встроенные пакеты crypto/hmac, crypto/sha256.

```go
import (
	"crypto/hmac"
	"crypto/sha256"
)
```

Пример генерации HMAC для номера карты:

```go
func ComputeHMAC(data string, secret []byte) string {
    h := hmac.New(sha256.New, secret)
    h.Write([]byte(data))
    return hex.EncodeToString(h.Sum(nil))
}
```

## Сдача проект

### Размещение кода на GitHub

Для сдачи нужно загрузить код в репозиторий и добавить в корень файл README.md с описанием:

- Как пользоваться сервисом.
- Какие команды поддерживаются.
- Как протестировать код.

### Подготовка к отправке проекта на проверку

Нужно проверить, что:

- Все функции проекта работают корректно и соответствуют ТЗ.
- Код оформлен и структурирован по стандартам.
- В репозитории есть описание проекта и документация по каждому модулю.
- Если используются внешние библиотеки, добавлены инструкции по их установке.

## Критерии оценивания

### Реализация слоя моделей (8 баллов)

- Определение структур данных (соответствие таблицам БД) — 2 балла.
- Сериализация/десериализация (теги JSON) — 2 балла.
- Базовая валидация полей (email, username) — 1 балл.
- Проверка уникальности (email, username) — 2 балла.
- Полная валидация всех полей — 1 балл.

### Реализация слоя репозиториев (9 баллов)

- Инкапсуляция SQL-запросов — 2 балла.
- Параметризованные запросы — 2 балла.
- Простейшая обработка ошибок БД — 1 балл.
- Управление транзакциями — 2 балла.
- Обработка сложных ошибок БД — 2 балла.

### Реализация слоя сервисов (20 баллов)

- Регистрация и аутентификация — 2 балла.
- Создание счетов, пополнение баланса — 3 балла.
- Переводы между счетами — 3 балла.
- Генерация карт (алгоритм Луна) — 2 балла.
- Кредиты: расчет аннуитетных платежей — 2 балла.
- Интеграция с SMTP (уведомления) — 2 балла.
- Интеграция с ЦБ РФ (SOAP) — 2 балла.
- Шедулер для списания платежей — 2 балла.
- Логирование через logrus — 2 балла.

### Реализация слоя обработчиков (12 баллов)

- Валидация входных данных — 2 балла.
- Формирование HTTP-ответов (JSON) — 2 балла.
- Вызов методов сервисов — 2 балла.
- Реализация всех эндпоинтов из ТЗ — 4 балла.
- Проверка прав доступа к ресурсам — 2 балла.

### Реализация маршрутизации (5 баллов)

- Публичные эндпоинты (/register, /login) — 2 балла.
- Защищенные эндпоинты (/accounts, /transfer и другие) — 3 балла.

### Реализация Middleware (6 баллов)

- Проверка JWT-токенов — 2 балла.
- Блокировка неавторизованных запросов — 2 балла.
- Добавление ID пользователя в контекст — 2 балла.

### Безопасность (7 баллов)

- Хеширование паролей (bcrypt) — 2 балла.
- Шифрование данных карт (PGP + HMAC) — 2 балла.
- Хеширование CVV (bcrypt) — 2 балла.
- Проверка прав доступа к счетам — 1 балл.

### База данных (2 балла)

- Создание минимальных таблиц — 2 балла.
