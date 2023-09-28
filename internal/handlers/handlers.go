package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Name        string `json:"taskTitle"`
	User        string `json:"taskDescription"`
	Description string `json:"taskAssignee"`
}

func PostNewTask(w http.ResponseWriter, r *http.Request) {
	var requests Request

	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	fmt.Println(requests, "fdvfasfasdf")
	w.WriteHeader(http.StatusCreated)
}
