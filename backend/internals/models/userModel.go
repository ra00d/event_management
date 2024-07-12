package models

type User struct {
	ID       string
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int
	IsActive bool `json:"is_active"`
}
