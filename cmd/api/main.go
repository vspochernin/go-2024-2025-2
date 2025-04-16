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

	"bank-system/config"
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

	// Инициализация роутера
	r := mux.NewRouter()

	// TODO: Добавить middleware и обработчики

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