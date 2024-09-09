package shared

type User struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	Id          int    `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      uint   `json:"user_id"`
}
