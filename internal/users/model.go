// User model definition
package users

import "Album/internal/tasks"

type User struct {
	Id       int          `json: "primaryKey" db:"id"`
	Name     string       `json: "name" binding:"required"`
	Username string       `json: "name" binding:"required"`
	Password string       `json: "name" binding:"required"`
	Tasks    []tasks.Task `gorm: "foreignKey:UserID"`
}
