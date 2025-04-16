package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"banksystem/config"
	"banksystem/internal/handlers"
	"banksystem/internal/middleware"
	"banksystem/internal/repositories"
	"banksystem/internal/services"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	// Загрузка конфигурации БД
	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatal("Ошибка загрузки конфигурации БД:", err)
	}

	// Подключение к БД
	db, err := sql.Open("postgres", dbConfig.GetDSN())
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	// Проверка подключения
	if err := db.Ping(); err != nil {
		log.Fatal("Ошибка проверки подключения к БД:", err)
	}
	fmt.Println("Успешное подключение к БД")

	// Инициализация репозиториев
	userRepo := repositories.NewUserRepository(db)
	accountRepo := repositories.NewAccountRepository(db)
	cardRepo := repositories.NewCardRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	creditRepo := repositories.NewCreditRepository(db)
	creditPaymentRepo := repositories.NewCreditPaymentRepository(db)

	// Инициализация сервисов
	authService := services.NewAuthService(userRepo)
	accountService := services.NewAccountService(accountRepo, transactionRepo)
	cardService := services.NewCardService(cardRepo, accountRepo)
	creditService := services.NewCreditService(creditRepo, accountRepo)
	creditPaymentService := services.NewCreditPaymentService(creditPaymentRepo, creditRepo, accountRepo)

	// Инициализация обработчиков
	authHandler := handlers.NewAuthHandler(authService)
	accountHandler := handlers.NewAccountHandler(accountService)
	cardHandler := handlers.NewCardHandler(cardService)
	creditHandler := handlers.NewCreditHandler(creditService)
	creditPaymentHandler := handlers.NewCreditPaymentHandler(creditPaymentService)

	// Инициализация роутера
	r := mux.NewRouter()

	// Публичные маршруты
	r.HandleFunc("/api/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/api/login", authHandler.Login).Methods("POST")

	// Защищенные маршруты
	protectedRouter := r.PathPrefix("/api").Subrouter()
	protectedRouter.Use(middleware.AuthMiddleware)

	// Маршруты для счетов
	protectedRouter.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	protectedRouter.HandleFunc("/accounts", accountHandler.GetUserAccounts).Methods("GET")
	protectedRouter.HandleFunc("/accounts/{id}/balance", accountHandler.GetBalance).Methods("GET")
	protectedRouter.HandleFunc("/accounts/{id}/deposit", accountHandler.Deposit).Methods("POST")
	protectedRouter.HandleFunc("/accounts/{id}/withdraw", accountHandler.Withdraw).Methods("POST")

	// Маршруты для карт
	protectedRouter.HandleFunc("/cards", cardHandler.CreateCard).Methods("POST")
	protectedRouter.HandleFunc("/cards", cardHandler.GetUserCards).Methods("GET")
	protectedRouter.HandleFunc("/cards/{id}", cardHandler.GetCard).Methods("GET")

	// Маршруты для кредитов
	protectedRouter.HandleFunc("/credits", creditHandler.CreateCredit).Methods("POST")
	protectedRouter.HandleFunc("/credits", creditHandler.GetUserCredits).Methods("GET")
	protectedRouter.HandleFunc("/credits/{id}", creditHandler.GetCredit).Methods("GET")

	// Маршруты для платежей по кредитам
	protectedRouter.HandleFunc("/credit-payments", creditPaymentHandler.CreatePayment).Methods("POST")
	protectedRouter.HandleFunc("/credit-payments/process", creditPaymentHandler.ProcessPayment).Methods("POST")
	protectedRouter.HandleFunc("/credit-payments", creditPaymentHandler.GetPaymentsByCreditID).Methods("GET")
	protectedRouter.HandleFunc("/credit-payments/pending", creditPaymentHandler.GetPendingPayments).Methods("GET")

	// Маршруты для переводов
	protectedRouter.HandleFunc("/transfer", accountHandler.Transfer).Methods("POST")

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Сервер запущен на порту %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
} 