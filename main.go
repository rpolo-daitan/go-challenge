package main

import (
	"fmt"
	"go-challenge/internal/repository"
	"log"
)

func main() {
	// Connect to database
	db, err := repository.Connect("172.17.0.2", "go_onb")
	if err != nil {
		log.Fatal(err)
	}

	// List all the tasks
	tasks, err := repository.GetAllTasks(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Tasks found: %v\n", tasks)

	// Get task by ID
	task, err := repository.GetTaskById(db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Task found for ID 2: %v\n", task)

	var testTask repository.Task

	// Create a new task
	testTask.Completed = true
	testTask.Name = "task insert test"
	ID, err := repository.AddTask(db, testTask)
	if err != nil {
		log.Fatal(err)
	}
	testTask.ID = ID
	fmt.Printf("Test task inserted. New task: %v\n", testTask)

	// Update a task
	testTask.Name = "task update test"
	testTask.Completed = false
	err = repository.UpdateTask(db, testTask)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test task updated. Updated task: %v\n", testTask)

	// List tasks by completion: completed tasks
	tasks, err = repository.GetAllTasksByCompletion(db, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Completed Tasks found: %v\n", tasks)

	// List tasks by completion: uncompleted tasks
	tasks, err = repository.GetAllTasksByCompletion(db, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Uncompleted Tasks found: %v\n", tasks)

	// Close database connection
	db.Close()
}
