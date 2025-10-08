package models

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUser(username string, email string) *User {
	return &User{
		Username: username,
		Email:    email,
	}
}
