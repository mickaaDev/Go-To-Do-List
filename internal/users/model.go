// User model definition
package users

import "Album/internal/shared"

type User struct {
	shared.User
	Tasks []shared.Task `gorm:"foreignKey:UserID"`
}
