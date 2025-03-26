package models

type (
	Guest struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Phone   string `json:"phone,omitempty"`
		Address string `json:"address"`
		Notes   any    `json:"notes"`
	}
)
