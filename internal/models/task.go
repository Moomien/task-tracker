package models

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"` //"todo" , "in-progress", "done"
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
