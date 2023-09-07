package models

type DBUser struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
