package model

type User struct {
	ID    int    `json:"id_user" db:"id"`
	Name  string `json:"name" db:"user_name"`
	Email string `json:"email" db:"email"`
}
