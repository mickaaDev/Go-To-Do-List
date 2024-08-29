// Task model definition
package tasks

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
