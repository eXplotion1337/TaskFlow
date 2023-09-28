package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Task struct {
	ID           int    `json:"taskID"`
	UID          string `json:"uid"`
	Name         string `json:"taskTitle"`
	NameProjects string `json:"NameProjects"`
	User         string `json:"taskAssignee"`
	Description  string `json:"taskDescription"`
	TimeStart    string `json:"timeStart"`
	TimeEnd      string `json:"TimeEnd"`
	Creator      string `json:"Creator"`
	Colum        string `json:"Colum"`
}

type MoveTask struct {
	Name   string `json:"Name"`
	Column string `json:"Column"`
}

type Project struct {
	NameProject   string `json:"NameProjects"`
	Description   string `json:"Description"`
	Collaborators string `json:"Collaborators"`
}

type Storage interface {
	AddNewTask(task Task) error
	AddNewProject(project Project) error
	SelectAllTasks() ([]Task, error)
	ToMoveTask(task MoveTask) error
}

type Postgres struct {
	db *sql.DB
}

func NewDatabaseStorage(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (db *Postgres) AddNewTask(task Task) error {
	insertQuery := `
		INSERT INTO tasks (name, project, "user", creator, time_start, time_end, colum_task, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	task.Colum = "Свободные"
	task.Creator = "Admin"
	_, err := db.db.Exec(insertQuery, task.Name, task.NameProjects, task.User, task.Creator, task.TimeStart, task.TimeEnd, task.Colum, task.Description)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (db *Postgres) AddNewProject(project Project) error {
	insertQuery := `
		INSERT INTO projects (name, collaborators, description)
		VALUES ($1, $2, $3)
	`

	_, err := db.db.Exec(insertQuery, project.NameProject, project.Collaborators, project.Description)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (db *Postgres) SelectAllTasks() ([]Task, error) {
	var result []Task

	selectQuery := `
		SELECT * FROM tasks 
	`

	rows, err := db.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.UID, &task.Name, &task.Description, &task.User, &task.TimeStart, &task.TimeEnd, &task.Creator, &task.Colum, &task.NameProjects); err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (db *Postgres) ToMoveTask(task MoveTask) error {
	query := `
        UPDATE tasks
        SET colum_task = $1
        WHERE name = $2 
    `

	_, err := db.db.Exec(query, task.Column, task.Name)
	if err != nil {
		log.Printf("ошибка изменения флага %s", err)
		return err
	}
	return nil
	return nil
}
