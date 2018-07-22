package models

import "github.com/go-ozzo/ozzo-validation"

// User represents a user record.
type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Role string `json:"role" db:"role"`
}

// Validate validates the User fields.
func (m User) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 100)),
	)
}
