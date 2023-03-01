package dto

import "time"

type UserDTO struct {
	FirstName string    `json:"first_name" binding:"required,min=2"`
	LastName  string    `json:"last_name" binding:"required,min=2"`
	Birthday  time.Time `json:"birthday" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Gender    string    `json:"gender" binding:"required,oneof=V M"`
	Password  string    `json:"password" binding:"required,min=6"`
}

type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
