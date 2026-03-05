package model

type User struct {
	ID    int    `json:"id_user"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
