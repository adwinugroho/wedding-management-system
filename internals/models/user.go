package models

import "time"

type (
	User struct {
		ID        string     `json:"id"`
		GoogleID  *string    `json:"google_id,omitempty"` // Nullable for local users
		Provider  string     `json:"provider"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		Password  *string    `json:"password,omitempty"` // Nullable for Google SSO users
		AvatarURL *string    `json:"avatar_url,omitempty"`
		Role      string     `json:"role"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		LastLogin *time.Time `json:"last_login,omitempty"`
	}
)
