package repository

import (
	"TaskFlow/internal/auth"
	"database/sql"
	"errors"
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
	ID            string `json:"ID"`
	NameProject   string `json:"NameProject"`
	Description   string `json:"Description"`
	Collaborators string `json:"Collaborators"`
	Token         string `json:"Token"`
}

type Storage interface {
	AddNewTask(task Task) error
	AddNewProject(project Project) error
	SelectAllTasks() ([]Task, error)
	ToMoveTask(task MoveTask) error
	SelectAllProjects() ([]Project, error)
	SelectTaskByProject(name string) ([]Task, error)
	CheckProdjectToExist(name string) (bool, error)
	GetUserByUsername(username string) (*auth.User, error)
	UserExists(username string) (bool, error)
	CreateUser(user auth.User) error
	Ping(dsn string)
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
	//project.ID = "1"
	insertQuery := `
		INSERT INTO projects (name, collaborators, description)
		VALUES ($1, $2, $3)
	`
	fmt.Println(project)
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
}

func (db *Postgres) SelectAllProjects() ([]Project, error) {
	var result []Project

	selectQuery := `
		SELECT * FROM projects 
	`

	rows, err := db.db.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.ID, &project.NameProject, &project.Description, &project.Collaborators, &project.Token); err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(result)
	return result, nil
}

func (db *Postgres) SelectTaskByProject(name string) ([]Task, error) {
	var result []Task
	query := `
		SELECT * FROM tasks
		WHERE project = $1
	`

	rows, err := db.db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.UID, &task.Name, &task.Description, &task.User, &task.TimeStart, &task.TimeEnd, &task.Creator, &task.Colum, &task.NameProjects)
		if err != nil {
			return nil, err
		}
		result = append(result, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (db *Postgres) CheckProdjectToExist(name string) (bool, error) {
	query := "SELECT COUNT(*) FROM projects WHERE name = $1"
	fmt.Println(name)
	// Выполняем запрос к базе данных.
	var count int
	err := db.db.QueryRow(query, name).Scan(&count)
	if err != nil {
		// Обработка ошибок при выполнении запроса.
		if errors.Is(err, sql.ErrNoRows) {
			// Если нет строк, значит проекта с таким именем нет.
			return false, nil
		}
		return false, fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}

	// Если count больше 0, значит проект существует.
	return count > 0, nil
}

func (db *Postgres) GetUserByUsername(username string) (*auth.User, error) {
	query := "SELECT id, login, password, token FROM users WHERE login = $1"
	row := db.db.QueryRow(query, username)

	var user auth.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Token)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *Postgres) UserExists(username string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)"
	var exists bool
	err := db.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (db *Postgres) CreateUser(user auth.User) error {
	query := "INSERT INTO users (login, password, token) VALUES ($1, $2, $3)"
	_, err := db.db.Exec(query, user.Username, user.Password, user.Token)
	return err
}

func (db *Postgres) Ping(dsn string) {
	err := db.db.Ping()
	if err != nil {
		fmt.Println(err)
		//http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("Apply")
	//return nil
}
