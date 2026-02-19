package main

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
	tasks []Task
}

var id int = 0

// add task
func (t *Tracker) Add(note string) {

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

}

func main() {
	t := Tracker{}
	t.Add("купить молоко")
	t.List("")
}
