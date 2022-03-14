package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Task entity.
type Task struct {
	ID        int64
	Name      string
	Completed bool
}

// Connect to a database.
func Connect(addr string, dbName string) (*sql.DB, error) {
	var db *sql.DB

	// Capture connection properties.
	cfg := mysql.Config{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   addr,
		DBName: dbName,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// List all the tasks.
func GetAllTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		return nil, fmt.Errorf("GetAllTasks %v", err)
	}
	return mapTasks(rows)
}

// Get task by ID
func GetTaskById(db *sql.DB, ID int64) (Task, error) {
	var task Task

	row := db.QueryRow("SELECT * FROM task WHERE ID=?", ID)
	if err := row.Scan(&task.ID, &task.Name, &task.Completed); err != nil {
		if err == sql.ErrNoRows {
			return task, fmt.Errorf("GetTaskById %d: no such task", ID)
		}
		return task, fmt.Errorf("GetTaskById %d: %v", ID, err)
	}
	return task, nil
}

// addTask adds the specified task to the database,
// returning the task ID of the new entry
func AddTask(db *sql.DB, task Task) (int64, error) {
	result, err := db.Exec("INSERT INTO task (Name, Completed) VALUES (?, ?)",
		task.Name, task.Completed)
	if err != nil {
		return 0, fmt.Errorf("AddTask: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddTask: %v", err)
	}
	return id, nil
}

// updateTask updates a specified task into the database.
func UpdateTask(db *sql.DB, task Task) error {
	_, err := db.Exec("UPDATE task SET Name=?, Completed=? WHERE ID=?",
		task.Name, task.Completed, task.ID)
	if err != nil {
		return fmt.Errorf("UpdateTask: %v", err)
	}
	return nil
}

// List tasks by completion.
func GetAllTasksByCompletion(db *sql.DB, Completion bool) ([]Task, error) {
	//var tasks []Task

	rows, err := db.Query("SELECT * FROM task WHERE Completed=?", Completion)
	if err != nil {
		return nil, fmt.Errorf("GetAllTasksByCompletion %v", err)
	}
	return mapTasks(rows)
}

// Maps task data from rows returned by a database query.
func mapTasks(rows *sql.Rows) ([]Task, error) {
	var tasks []Task

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Completed); err != nil {
			return nil, fmt.Errorf("mapTasks %v", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("mapTasks %v", err)
	}
	return tasks, nil
}
