package app

import (
	"TaskFlow/internal/config"
	"TaskFlow/internal/handlers"
	my_middleware "TaskFlow/internal/middleware"
	"TaskFlow/internal/repository"
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strings"
)

func Run(config *config.Config, storage repository.Storage) error {
	router := chi.NewRouter()

	// Middleware для логирования запросов
	router.Use(middleware.Logger)
	router.Use(my_middleware.AuthMiddleware)
	router.Use(my_middleware.PromMiddleware)

	router.Get("/prom/metrics", func(w http.ResponseWriter, r *http.Request) {
		//promhttp.Handler()
		promhttp.Handler().ServeHTTP(w, r)
	})

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
		handlers.PostNewTask(w, r, storage)
	})

	router.Get("/createNewProject", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/createNewProject/createNewProject.html")
	})

	router.Post("/createNewProject", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostNewProject(w, r, storage)
	})

	router.Post("/api/alltask", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostAllTasks(w, r, storage)
	})

	router.Post("/api/movetask", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostMoveTask(w, r, storage)
	})

	router.Get("/Projects", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/Projects/projects.html")
	})

	router.Post("/api/allprojects", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostAllProject(w, r, storage)
	})

	router.Get("/dashboard/{id}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/dashboard/task_for_project/task_for_project.html")
	})

	router.Post("/api/dashboard/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostSelectTaskByProject(w, r, storage)
	})

	router.Post("/api/checkProdject/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostCheckProjectExist(w, r, storage)
	})

	router.Post("/auth/sing-in", func(w http.ResponseWriter, r *http.Request) {
		handlers.SingIn(w, r, storage)
	})

	router.Post("/auth/sing-up", func(w http.ResponseWriter, r *http.Request) {
		handlers.SingUp(w, r, storage)
	})

	router.Get("/sing-in", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/auth/sing-in.html")
	})

	router.Get("/sing-up", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/auth/sing-up.html")
	})

	router.Get("/ping/ping", func(w http.ResponseWriter, r *http.Request) {
		handlers.Ping(w, r, storage)
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
		DataBase: &config.DataBase{
			DSN: GetEnvAsStr("DSN", "postgresql://postgres:123123@localhost:5432/postgres?sslmode=disable"),
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

func InitStorage(config *config.Config) (repository.Storage, error) {
	db, err := sql.Open("postgres", config.DataBase.DSN)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	storage := repository.NewDatabaseStorage(db)

	return storage, nil
}
