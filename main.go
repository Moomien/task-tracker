package main

import (
	"fmt"
	"strconv"
	"time"
)

type taskTracker interface {
	Add(note string)
	Update(id string, note string)
	Delete(id string)
	MarkInProgress(id string)
	MarkDone(id string)
	List(status string)
}

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"` //"todo" , "in-progress", "done"
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Tracker struct {
	tasks   []Task
	storage Storage
}

type Storage interface {
	SaveTasks(tasks []Task) error
	LoadTasks() ([]Task, error)
}

type JSONStorage struct {
	filename string
}

func NewTracker(storage Storage) *Tracker {
	tasks, _ := storage.LoadTasks()
	return &Tracker{
		tasks:   tasks,
		storage: storage,
	}
}

func NewJSONStorage(filename string) *JSONStorage {
	return &JSONStorage{filename: filename}
}

var id int = 0

// add task
func (t *Tracker) Add(note string) {
	id++
	time := time.Now()
	t.tasks = append(t.tasks, Task{
		ID:          strconv.Itoa(id),
		Description: note,
		Status:      "todo",
		CreatedAt:   time.Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Format("2006-01-02 15:04:05"),
	})
}

// update task
func (t *Tracker) Update(id string, note string) {

}

// delete task
func (t *Tracker) Delete(id string) {

}

// mark progress "in-progress"
func (t *Tracker) MarkInProgress(id string) {

}

// mark progress "done"
func (t *Tracker) MarkDone(id string) {

}

// listing tasks by status
func (t *Tracker) List(status string) {
	switch status {
	case "":
		for _, v := range t.tasks {
			fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s",
				v.ID,
				v.Description,
				v.Status,
				v.CreatedAt,
				v.UpdatedAt)
		}
	case "todo":
		for _, v := range t.tasks {
			if v.Status == "todo" {
				fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s",
					v.ID,
					v.Description,
					v.Status,
					v.CreatedAt,
					v.UpdatedAt)
			}
		}
	case "in-progress":
		for _, v := range t.tasks {
			if v.Status == "in-progress" {
				fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s",
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
				fmt.Printf("ID: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s",
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

func main() {
	t := Tracker{}
	t.Add("купить молоко")
	t.List("")
}
