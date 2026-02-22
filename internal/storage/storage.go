package storage

import (
	"encoding/json"
	"os"
	"task-tracker/internal/models"
)

type Storage interface {
	SaveTasks(tasks []models.Task) error
	LoadTasks() ([]models.Task, error)
}

type JSONStorage struct {
	filename string
}

// new json storage instance
func NewJSONStorage(filename string) *JSONStorage {
	return &JSONStorage{filename: filename + ".json"}
}

// save tasks to json storage
func (js *JSONStorage) SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(js.filename, data, 0644)
}

// load tasks from json storage
func (js *JSONStorage) LoadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(js.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []models.Task{}, nil
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}
