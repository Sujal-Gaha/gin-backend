package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"user_id"`
}
