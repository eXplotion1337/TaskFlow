package handlers

import (
	"TaskFlow/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
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

	//body, _ := io.ReadAll(r.Body)
	//fmt.Println(string(body))

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
