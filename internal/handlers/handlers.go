package handlers

import (
	"TaskFlow/internal/auth"
	"TaskFlow/internal/repository"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
)

func PostNewTask(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	var requests repository.Task

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	err := storage.AddNewTask(requests)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PostNewProject(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	var requests repository.Project

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	err := storage.AddNewProject(requests)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func PostAllTasks(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	tasks, err := storage.SelectAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	taskJSON, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Установите заголовок Content-Type на application/json.
	w.Header().Set("Content-Type", "application/json")

	// Установите статус HTTP-ответа на 200 OK и отправьте JSON.
	w.WriteHeader(http.StatusOK)
	w.Write(taskJSON)

}

func PostMoveTask(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	var requestData repository.MoveTask

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Println(err, requestData)
		return
	}

	err = storage.ToMoveTask(requestData)
	if err != nil {
		http.Error(w, "ауйайай", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func PostAllProject(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	projects, err := storage.SelectAllProjects()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	projectsJSON, err := json.Marshal(projects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Установите заголовок Content-Type на application/json.
	w.Header().Set("Content-Type", "application/json")

	// Установите статус HTTP-ответа на 200 OK и отправьте JSON.
	w.WriteHeader(http.StatusOK)
	w.Write(projectsJSON)
}

func PostSelectTaskByProject(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	tasks, err := storage.SelectTaskByProject(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Установите заголовок Content-Type на application/json.
	w.Header().Set("Content-Type", "application/json")

	// Установите статус HTTP-ответа на 200 OK и отправьте JSON.
	w.WriteHeader(http.StatusOK)
	w.Write(tasksJSON)

}

func PostCheckProjectExist(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	id := chi.URLParam(r, "id")
	b, err := storage.CheckProdjectToExist(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if b {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func SingIn(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	// Получаем данные из тела запроса
	var user auth.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	// Получаем хэшированный пароль из базы данных по имени пользователя
	storedUser, err := storage.GetUserByUsername(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid username or password"))
		return
	}

	// Проверяем, соответствует ли введенный пароль хэшированному паролю в базе данных
	if err := auth.CheckPassword(user.Password, storedUser.Password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid username or password"))
		return
	}

	// Генерируем токен и отправляем его клиенту
	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error generating token"))
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"token": token}
	json.NewEncoder(w).Encode(response)
}

func SingUp(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	// Получаем данные из тела запроса
	var user auth.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}

	user1, _ := storage.UserExists(user.Username)
	// Проверяем, не существует ли уже пользователь с таким именем
	if user1 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("User with this username already exists"))
		return
	}

	// Хэшируем пароль перед сохранением в базу данных
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error hashing password"))
		return
	}

	// Сохраняем пользователя в базу данных
	newUser := auth.User{Username: user.Username, Password: hashedPassword}
	if err := storage.CreateUser(newUser); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating user"))
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func Ping(w http.ResponseWriter, r *http.Request, storage repository.Storage) {
	dsn := os.Getenv("DSN")
	storage.Ping(dsn)
}
