package service

import (
	"fmt"
	"strconv"
	"strings"
	"task-tracker/internal/models"
	"task-tracker/internal/storage"
	"time"
)

type taskTracker interface {
	Add(note string)
	Update(id string, note string)
	Delete(id string)
	Mark(id string, status string)
	List(status string)
}

type Tracker struct {
	tasks   []models.Task
	storage storage.Storage
}

// new tracker instance
func NewTracker(storage storage.Storage) *Tracker {
	tasks, _ := storage.LoadTasks()
	return &Tracker{
		tasks:   tasks,
		storage: storage,
	}
}

// add task
func (t *Tracker) Add(note string) {
	tasks, err := t.storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		return
	}
	//generate new id based on max existing
	maxID := 0
	for _, task := range tasks {
		if id, err := strconv.Atoi(task.ID); err == nil && id > maxID {
			maxID = id
		}
	}

	newTask := models.Task{
		ID:          strconv.Itoa(maxID + 1),
		Description: note,
		Status:      "todo",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	//save tasks to storage
	tasks = append(tasks, newTask)
	if err := t.storage.SaveTasks(tasks); err != nil {
		fmt.Println("error while saving tasks to storage: ", err)
		return
	}
}

// update task
func (t *Tracker) Update(id string, note string) {
	tasks, err := t.storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		return
	}
	for i, v := range tasks {
		if v.ID == id {
			tasks[i] = models.Task{
				ID:          id,
				Description: note,
				Status:      "todo",
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			}
			if err := t.storage.SaveTasks(tasks); err != nil {
				fmt.Println("error while saving tasks to storage: ", err)
				return
			}
			fmt.Printf("Task %s updated!\n", id)
			break
		}
	}
}

// delete task
func (t *Tracker) Delete(id string) {
	tasks, err := t.storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		return
	}

	for i, v := range tasks {
		if v.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			if err := t.storage.SaveTasks(tasks); err != nil {
				fmt.Println("error while saving tasks to storage: ", err)
				return
			}
			fmt.Printf("Task %s deleted\n", id)
			break
		}
	}
}

// marg progress ("done", "in-progress")
func (t *Tracker) Mark(id string, status string) {
	tasks, err := t.storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		return
	}

	for i, v := range tasks {
		if v.ID == id {
			switch strings.ToLower(status) {
			case "done":
				tasks[i] = models.Task{
					ID:          id,
					Description: v.Description,
					Status:      "Done",
					CreatedAt:   v.CreatedAt,
					UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				}
			case "in-progress":
				tasks[i] = models.Task{
					ID:          id,
					Description: v.Description,
					Status:      "In-Progress",
					CreatedAt:   v.CreatedAt,
					UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				}
			default:
				fmt.Println("Wrong argument!")
				return
			}
			if err := t.storage.SaveTasks(tasks); err != nil {
				fmt.Println("error while saving tasks to storage: ", err)
				return
			}
			fmt.Printf("Task %s updated!\n", id)
			break
		}
	}
}

// listing tasks by status
func (t *Tracker) List(status string) {
	switch strings.ToLower(status) {
	case "all":
		for _, v := range t.tasks {
			fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				v.ID,
				v.Description,
				v.Status,
				v.CreatedAt,
				v.UpdatedAt)
		}
	case "todo":
		for _, v := range t.tasks {
			if v.Status == "todo" {
				fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
					v.ID,
					v.Description,
					v.Status,
					v.CreatedAt,
					v.UpdatedAt)
			}
		}
	case "in-progress":
		for _, v := range t.tasks {
			if v.Status == "In-Progress" {
				fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
					v.ID,
					v.Description,
					v.Status,
					v.CreatedAt,
					v.UpdatedAt)
			}
		}
	case "done":
		for _, v := range t.tasks {
			if v.Status == "done" {
				fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
					v.ID,
					v.Description,
					v.Status,
					v.CreatedAt,
					v.UpdatedAt)
			}
		}
	default:
		fmt.Println("Wrong argument")
	}
}
