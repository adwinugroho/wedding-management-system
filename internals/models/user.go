package models

import "time"

type (
	User struct {
		ID        string    `json:"id"`
		Role      string    `json:"role"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
