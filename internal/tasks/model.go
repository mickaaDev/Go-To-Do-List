// Task model definition
package tasks

type Task struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
}
