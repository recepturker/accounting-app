package models

type Partner struct {
	ID          uint   `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Email       string `db:"email" json:"email"`
	Phone       string `db:"phone" json:"phone"`
	Address     string `db:"address" json:"address"`
	Logo        string `db:"logo" json:"logo"`
	Aproved     bool   `db:"aproved" json:"aproved"`
	DefaultUser uint   `db:"default_user" json:"default_user"`
}
