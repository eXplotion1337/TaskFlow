package app

import (
	"TaskFlow/internal/config"
	"TaskFlow/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"strings"
)

func Run(config *config.Config) error {
	router := chi.NewRouter()

	// Middleware для логирования запросов
	router.Use(middleware.Logger)

	// Обработчики для страниц
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/main/main.html")
	})

	router.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/dashboard/dashboard.html")
	})

	router.Get("/createNewTask", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/createNewTask/create-task.html")
	})

	router.Post("/createNewTask", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostNewTask(w, r)
	})

	// Обработчики для статических файлов (стили и скрипты)
	fs := http.FileServer(http.Dir("./web"))
	router.Handle("/styles/*", fs)
	router.Handle("/scripts/*", fs)

	log.Printf("Сервер запущен на порту  %s", config.Http.Port)
	log.Printf("Сервер запущен на адресе  %s", config.Http.Host)

	// Запуск сервера на порту 8080
	err := http.ListenAndServe(config.Http.Port, router)
	if err != nil {
		return err
	}

	return nil
}

func InitConfig() (*config.Config, error) {
	Config := config.Config{
		Http: &config.Http{
			Port: GetEnvAsStr("HTTP_PORT", ":8080"),
			Host: GetEnvAsStr("HTTP_HOST", "localhost"),
		},
	}
	return &Config, nil
}

func GetEnvAsStr(name string, defaultValue string) string {
	valStr := os.Getenv(name)
	if strings.TrimSpace(valStr) == "" {
		return defaultValue
	}

	return valStr
}
