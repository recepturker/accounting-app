package models

type User struct {
	BaseModel
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Active   bool   `json:"active" db:"active"`
}
